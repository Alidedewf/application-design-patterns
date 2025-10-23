package main

import (
	"fmt"
	"time"
)

type IMediator interface {
	Register(user IColleague)
	Unregister(user IColleague)
	SendMessage(message string, sender IColleague) error
	SendPrivateMessage(message string, sender IColleague, receiverName string) error

	CreateGroup(groupName string) error
	JoinGroup(groupName string, user IColleague) error
	SendMessageToGroup(groupName string, message string, sender IColleague) error
}

type IColleague interface {
	ReceiveMessage(message string)
	GetName() string
}

type User struct {
	mediator IMediator
	name     string
}

func NewUser(mediator IMediator, name string) *User {
	return &User{
		mediator: mediator,
		name:     name,
	}
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) ReceiveMessage(message string) {
	fmt.Printf("%s получил сообщение: %s\n", u.name, message)
}

func (u *User) Send(message string) {
	fmt.Printf("%s отправляет (всем): %s\n", u.name, message)

	err := u.mediator.SendMessage(message, u)

	if err != nil {
		fmt.Printf("[ошибка] %s: %s\n", u.name, err)
	}
}

func (u *User) SendPrivate(message string, receiverName string) {
	fmt.Printf("%s отправляет (лично %s): %s\n", u.name, receiverName, message)
	err := u.mediator.SendPrivateMessage(message, u, receiverName)
	if err != nil {
		fmt.Printf("[ошибка] %s: %s\n", u.name, err)
	}
}

func (u *User) SendToGroup(groupName string, message string) {
	fmt.Printf("%s отправляет (группа %s): %s\n", u.name, groupName, message)
	err := u.mediator.SendMessageToGroup(groupName, message, u)
	if err != nil {
		fmt.Printf("[ошибка] %s: %s\n", u.name, err)
	}
}

type ChatMediator struct {
	users map[string]IColleague

	groups map[string]map[string]IColleague

	messageLog []string
}

func NewChatMediator() *ChatMediator {
	return &ChatMediator{
		users:      make(map[string]IColleague),
		groups:     make(map[string]map[string]IColleague),
		messageLog: make([]string, 0),
	}
}

func (m *ChatMediator) log(entry string) {
	logEntry := fmt.Sprintf("[%s] %s", time.Now().Format("15:04:05"), entry)
	m.messageLog = append(m.messageLog, logEntry)
	fmt.Printf("[лог] %s\n", entry)
}

func (m *ChatMediator) isRegistered(user IColleague) bool {
	_, ok := m.users[user.GetName()]
	return ok
}

func (m *ChatMediator) Register(user IColleague) {
	m.users[user.GetName()] = user
	m.log(fmt.Sprintf("участник '%s' зарегистрирован в системе.", user.GetName()))
}

func (m *ChatMediator) Unregister(user IColleague) {
	userName := user.GetName()
	if _, ok := m.users[userName]; ok {
		delete(m.users, userName)

		for groupName, members := range m.groups {
			if _, ok := members[userName]; ok {
				delete(m.groups[groupName], userName)
			}
		}
		m.log(fmt.Sprintf("участник '%s' отключен от системы.", userName))
	}
}

func (m *ChatMediator) SendMessage(message string, sender IColleague) error {
	if !m.isRegistered(sender) {
		return fmt.Errorf("вы не зарегистрированы в чате")
	}

	fullMessage := fmt.Sprintf("[%s]: %s", sender.GetName(), message)
	m.log(fmt.Sprintf("сообщение (всем) от '%s': %s", sender.GetName(), message))

	for _, user := range m.users {
		if user.GetName() != sender.GetName() {
			user.ReceiveMessage(fullMessage)
		}
	}
	return nil
}

func (m *ChatMediator) SendPrivateMessage(message string, sender IColleague, receiverName string) error {
	if !m.isRegistered(sender) {
		return fmt.Errorf("вы не зарегистрированы в чате")
	}

	receiver, ok := m.users[receiverName]
	if !ok {
		return fmt.Errorf("участник '%s' не найден", receiverName)
	}

	fullMessage := fmt.Sprintf("[%s -> вам]: %s", sender.GetName(), message)
	m.log(fmt.Sprintf("сообщение (личное) от '%s' для '%s'", sender.GetName(), receiverName))

	receiver.ReceiveMessage(fullMessage)
	sender.ReceiveMessage(fmt.Sprintf("[вы -> %s]: %s", receiverName, message))
	return nil
}

func (m *ChatMediator) CreateGroup(groupName string) error {
	if _, ok := m.groups[groupName]; ok {
		return fmt.Errorf("группа '%s' уже существует", groupName)
	}
	m.groups[groupName] = make(map[string]IColleague)
	m.log(fmt.Sprintf("создана группа '%s'", groupName))
	return nil
}

func (m *ChatMediator) JoinGroup(groupName string, user IColleague) error {
	if !m.isRegistered(user) {
		return fmt.Errorf("сначала зарегистрируйтесь в системе")
	}
	group, ok := m.groups[groupName]
	if !ok {
		return fmt.Errorf("группа '%s' не найдена", groupName)
	}

	userName := user.GetName()
	if _, ok := group[userName]; ok {
		return nil
	}

	m.groups[groupName][userName] = user
	m.log(fmt.Sprintf("участник '%s' вступил в группу '%s'", userName, groupName))
	return nil
}

func (m *ChatMediator) SendMessageToGroup(groupName string, message string, sender IColleague) error {
	if !m.isRegistered(sender) {
		return fmt.Errorf("вы не зарегистрированы в чате")
	}
	group, ok := m.groups[groupName]
	if !ok {
		return fmt.Errorf("группа '%s' не найдена", groupName)
	}

	senderName := sender.GetName()
	if _, ok := group[senderName]; !ok {
		return fmt.Errorf("вы не состоите в группе '%s'", groupName)
	}

	fullMessage := fmt.Sprintf("[%s@%s]: %s", senderName, groupName, message)
	m.log(fmt.Sprintf("сообщение (группа %s) от '%s'", groupName, senderName))

	for name, user := range group {
		if name != senderName {
			user.ReceiveMessage(fullMessage)
		}
	}
	return nil
}

func main() {
	chatMediator := NewChatMediator()

	user1 := NewUser(chatMediator, "алихан")
	user2 := NewUser(chatMediator, "андрей")
	user3 := NewUser(chatMediator, "стас")
	user4 := NewUser(chatMediator, "влад")

	chatMediator.Register(user1)
	chatMediator.Register(user2)
	chatMediator.Register(user3)

	fmt.Println("\n--- тест 1: общий чат ---")
	user1.Send("привет всем!")

	replyMsg := fmt.Sprintf("привет, %s!", user1.GetName())
	user2.Send(replyMsg)

	fmt.Println("\n--- тест 2: приватные сообщения ---")

	privateMsg := fmt.Sprintf("%s, как дела?", user2.GetName())
	user3.SendPrivate(privateMsg, user2.GetName())

	user1.SendPrivate("ты не существуешь", "!")

	fmt.Println("\n--- тест 3: отключение участника ---")
	chatMediator.Unregister(user2)

	user1.Send(fmt.Sprintf("%s, ты куда ушел? тебя удалили.", user2.GetName()))

	user3.SendPrivate(fmt.Sprintf("%s, ты еще здесь?", user2.GetName()), user2.GetName())

	fmt.Println("\n--- тест 4: отправка без регистрации ---")
	user4.Send("меня кто-нибудь слышит?")

	fmt.Println("\n--- тест 5: групповые чаты ---")
	chatMediator.CreateGroup("разработчики")
	chatMediator.CreateGroup("менеджеры")

	chatMediator.JoinGroup("разработчики", user1)
	chatMediator.JoinGroup("разработчики", user3)
	chatMediator.JoinGroup("менеджеры", user3)

	err := chatMediator.JoinGroup("менеджеры", user2)
	if err != nil {
		fmt.Printf("[ошибка] %s: %s\n", user2.GetName(), err)
	}

	user1.SendToGroup("разработчики", "нужно исправить баг")
	user3.SendToGroup("менеджеры", "отчет готов")

	user1.SendToGroup("менеджеры", "я не менеджер, но... ")
	user3.SendToGroup("разработчики", "баг принят")

	fmt.Println("\n--- тест 6: логирование ---")
	fmt.Println("логирование выводилось в реальном времени с префиксом.")
}

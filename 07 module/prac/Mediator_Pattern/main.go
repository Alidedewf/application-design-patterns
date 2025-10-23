package main

import (
	"fmt"
)

type IUser interface {
	ReceiveMessage(message string)
	GetName() string
}

type IMediator interface {
	RegisterUser(user IUser)
	SetAdmin(userName string)
	JoinChannel(channelName string, user IUser) error
	LeaveChannel(channelName string, user IUser) error
	BroadcastToChannel(channelName, message string, sender IUser) error
	SendPrivateMessage(receiverName, message string, sender IUser) error
	SendCrossChannelMessage(targetChannel, message string, sender IUser) error
	BanUser(adminName, targetUserName, channelName string) error
}

type User struct {
	name     string
	mediator IMediator
}

func NewUser(name string, mediator IMediator) *User {
	u := &User{name: name, mediator: mediator}
	mediator.RegisterUser(u)
	return u
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) ReceiveMessage(message string) {
	fmt.Printf("%s получил: %s\n", u.name, message)
}

func (u *User) Join(channelName string) {
	fmt.Printf("%s пытается войти в канал '%s'\n", u.name, channelName)
	if err := u.mediator.JoinChannel(channelName, u); err != nil {
		u.ReceiveMessage(fmt.Sprintf("[ошибка] %s", err.Error()))
	}
}

func (u *User) Leave(channelName string) {
	fmt.Printf("%s покидает канал '%s'\n", u.name, channelName)
	if err := u.mediator.LeaveChannel(channelName, u); err != nil {
		u.ReceiveMessage(fmt.Sprintf("[ошибка] %s", err.Error()))
	}
}

func (u *User) Send(channelName, message string) {
	fmt.Printf("%s -> [%s]: %s\n", u.name, channelName, message)
	if err := u.mediator.BroadcastToChannel(channelName, message, u); err != nil {
		u.ReceiveMessage(fmt.Sprintf("[ошибка] %s", err.Error()))
	}
}

func (u *User) SendPrivate(receiverName, message string) {
	fmt.Printf("%s -> (лично %s): %s\n", u.name, receiverName, message)
	if err := u.mediator.SendPrivateMessage(receiverName, message, u); err != nil {
		u.ReceiveMessage(fmt.Sprintf("[ошибка] %s", err.Error()))
	}
}

func (u *User) SendCrossChannel(targetChannel, message string) {
	fmt.Printf("%s -> (в канал %s): %s\n", u.name, targetChannel, message)
	if err := u.mediator.SendCrossChannelMessage(targetChannel, message, u); err != nil {
		u.ReceiveMessage(fmt.Sprintf("[ошибка] %s", err.Error()))
	}
}

func (u *User) Ban(targetUserName, channelName string) {
	fmt.Printf("%s банит %s в канале %s\n", u.name, targetUserName, channelName)
	if err := u.mediator.BanUser(u.name, targetUserName, channelName); err != nil {
		u.ReceiveMessage(fmt.Sprintf("[ошибка] %s", err.Error()))
	}
}

type Channel struct {
	name    string
	members map[string]IUser
}

func NewChannel(name string) *Channel {
	return &Channel{
		name:    name,
		members: make(map[string]IUser),
	}
}

func (c *Channel) Broadcast(message string, sender IUser) {
	fullMsg := fmt.Sprintf("[%s@%s] %s: %s", sender.GetName(), c.name, sender.GetName(), message)
	for _, member := range c.members {
		if member.GetName() != sender.GetName() {
			member.ReceiveMessage(fullMsg)
		}
	}
}

func (c *Channel) Notify(message string) {
	fullMsg := fmt.Sprintf("[%s] (система): %s", c.name, message)
	for _, member := range c.members {
		member.ReceiveMessage(fullMsg)
	}
}

func (c *Channel) AddMember(user IUser) {
	c.members[user.GetName()] = user
	c.Notify(fmt.Sprintf("%s присоединился.", user.GetName()))
}

func (c *Channel) RemoveMember(user IUser) {
	delete(c.members, user.GetName())
	c.Notify(fmt.Sprintf("%s покинул.", user.GetName()))
}

type ChannelMediator struct {
	channels map[string]*Channel
	// Исправлено: убраны невидимые символы
	users    map[string]IUser
	banned   map[string]map[string]bool
	admins   map[string]bool
}

func NewChannelMediator() *ChannelMediator {
	// Исправлено: убраны невидимые символы
	return &ChannelMediator{
		channels: make(map[string]*Channel),
		users:    make(map[string]IUser),
		banned:   make(map[string]map[string]bool),
		admins:   make(map[string]bool),
	}
}

func (m *ChannelMediator) getOrCreateChannel(name string) *Channel {
	if channel, ok := m.channels[name]; ok {
		return channel
	}
	fmt.Printf("(система) создан новый канал: %s\n", name)
	newChannel := NewChannel(name)
	m.channels[name] = newChannel
	m.banned[name] = make(map[string]bool)
	return newChannel
}

func (m *ChannelMediator) isBanned(channelName, userName string) bool {
	if channelBans, ok := m.banned[channelName]; ok {
		if isBanned, ok := channelBans[userName]; ok && isBanned {
			return true
		}
	}
	return false
}

func (m *ChannelMediator) isAdmin(userName string) bool {
	return m.admins[userName]
}

func (m *ChannelMediator) RegisterUser(user IUser) {
	m.users[user.GetName()] = user
	fmt.Printf("%s зарегистрирован\n", user.GetName())
}

func (m *ChannelMediator) SetAdmin(userName string) {
	if _, ok := m.users[userName]; ok {
		m.admins[userName] = true
		fmt.Printf("%s назначен администратором\n", userName)
	}
}

func (m *ChannelMediator) JoinChannel(channelName string, user IUser) error {
	if m.isBanned(channelName, user.GetName()) {
		return fmt.Errorf("вы забанены в канале '%s'", channelName)
	}
	channel := m.getOrCreateChannel(channelName)
	channel.AddMember(user)
	return nil
}

func (m *ChannelMediator) LeaveChannel(channelName string, user IUser) error {
	channel, ok := m.channels[channelName]
	if !ok {
		return fmt.Errorf("канал '%s' не существует", channelName)
	}
	channel.RemoveMember(user)
	return nil
}

func (m *ChannelMediator) BroadcastToChannel(channelName, message string, sender IUser) error {
	channel, ok := m.channels[channelName]
	if !ok {
		return fmt.Errorf("канал '%s' не существует", channelName)
	}

	if _, ok := channel.members[sender.GetName()]; !ok {
		return fmt.Errorf("вы не состоите в канале '%s'", channelName)
	}

	if m.isBanned(channelName, sender.GetName()) {
		return fmt.Errorf("вы забанены в канале '%s'", channelName)
	}

	channel.Broadcast(message, sender)
	return nil
}

func (m *ChannelMediator) SendPrivateMessage(receiverName, message string, sender IUser) error {
	receiver, ok := m.users[receiverName]
	if !ok {
		return fmt.Errorf("пользователь '%s' не найден", receiverName)
	}

	fullMsg := fmt.Sprintf("(лично) %s: %s", sender.GetName(), message)
	receiver.ReceiveMessage(fullMsg)
	return nil
}

func (m *ChannelMediator) SendCrossChannelMessage(targetChannel, message string, sender IUser) error {
	channel, ok := m.channels[targetChannel]
	if !ok {
		return fmt.Errorf("канал '%s' не существует", targetChannel)
	}

	if m.isBanned(targetChannel, sender.GetName()) {
		return fmt.Errorf("вы забанены в канале '%s'", targetChannel)
	}

	fullMsg := fmt.Sprintf("[%s->%s] %s: %s", sender.GetName(), targetChannel, sender.GetName(), message)
	for _, member := range channel.members {
		member.ReceiveMessage(fullMsg)
	}
	return nil
}

func (m *ChannelMediator) BanUser(adminName, targetUserName, channelName string) error {
	if !m.isAdmin(adminName) {
		return fmt.Errorf("у вас нет прав администратора")
	}

	if _, ok := m.users[targetUserName]; !ok {
		return fmt.Errorf("пользователь '%s' не найден", targetUserName)
	}

	channel := m.getOrCreateChannel(channelName)

	m.banned[channelName][targetUserName] = true

	channel.Notify(fmt.Sprintf("%s был забанен админом %s.", targetUserName, adminName))

	if targetUser, ok := channel.members[targetUserName]; ok {
		channel.RemoveMember(targetUser)
	}

	return nil
}

func main() {
	mediator := NewChannelMediator()

	user1 := NewUser("Алихан", mediator)
	user2 := NewUser("Андрей", mediator)
	user3 := NewUser("Стас", mediator)
	user4 := NewUser("Влад", mediator)

	mediator.SetAdmin(user1.GetName())

	fmt.Println("\n--- тест 1: каналы и уведомления ---")
	user1.Join("общий")
	user2.Join("общий")
	user3.Join("игры")

	reply1 := fmt.Sprintf("привет, %s!", user2.GetName())
	user1.Send("общий", reply1)

	reply2 := fmt.Sprintf("привет, %s!", user1.GetName())
	user2.Send("общий", reply2)

	fmt.Println("\n--- тест 2: приватные сообщения ---")
	user1.SendPrivate(user3.GetName(), "секретное сообщение")

	fmt.Println("\n--- тест 3: выход и ошибки ---")
	user2.Leave("общий")
	user1.Send("общий", fmt.Sprintf("%s ушел", user2.GetName()))
	user2.Send("общий", "я не могу отправить")

	fmt.Println("\n--- тест 4: админ (бан) ---")
	user1.Ban(user4.GetName(), "игры")
	user4.Join("игры")
	user4.Send("игры", "я не могу отправить")

	fmt.Println("\n--- тест 5: авто-создание канала ---")
	user3.Join("музыка")
	user3.Send("музыка", "теперь я могу отправить")

	fmt.Println("\n--- тест 6: кросс-канал ---")
	user1.SendCrossChannel("музыка", "привет из другого канала")
}
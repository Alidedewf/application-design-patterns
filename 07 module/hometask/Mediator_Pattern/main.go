package main

import (
	"fmt"
)

type IMediator interface {
	RegisterUser(user *User)
	UnregisterUser(user *User) 
	SendMessage(message string, sender *User)
	SendPrivateMessage(message string, sender *User, receiverName string) 
}

type User struct {
	mediator IMediator 
	Name     string
}

func NewUser(name string) *User {
	return &User{Name: name}
}

func (u *User) JoinChat(chat IMediator) {
	u.mediator = chat
	u.mediator.RegisterUser(u)
}

func (u *User) LeaveChat() {
	if u.mediator == nil {
		return 
	}
	u.mediator.UnregisterUser(u)
	u.mediator = nil
}

func (u *User) Send(message string) {
	if u.mediator == nil {
		fmt.Printf("Ошибка: %s не может отправить сообщение. Сначала войдите в чат.\n", u.Name)
		return
	}
	u.mediator.SendMessage(message, u)
}

func (u *User) SendPrivate(message string, receiverName string) {
	if u.mediator == nil {
		fmt.Printf("Ошибка: %s не может отправить сообщение. Сначала войдите в чат.\n", u.Name)
		return
	}
	u.mediator.SendPrivateMessage(message, u, receiverName)
}

func (u *User) Receive(message string) {
	fmt.Printf("[%s's console]: %s\n", u.Name, message)
}

type ChatRoom struct {
	users map[string]*User
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		users: make(map[string]*User),
	}
}

func (cr *ChatRoom) broadcastMessage(message string, sender *User, includeSender bool) {
	for _, user := range cr.users {
		if user != sender || includeSender {
			user.Receive(message)
		}
	}
}

func (cr *ChatRoom) RegisterUser(user *User) {
	cr.users[user.Name] = user
	cr.broadcastMessage(fmt.Sprintf("--- %s присоединился к чату ---", user.Name), user, true)
}

func (cr *ChatRoom) UnregisterUser(user *User) {
	if _, ok := cr.users[user.Name]; ok {
		delete(cr.users, user.Name)
		cr.broadcastMessage(fmt.Sprintf("--- %s покинул чат ---", user.Name), user, true)
	}
}

func (cr *ChatRoom) SendMessage(message string, sender *User) {
	cr.broadcastMessage(fmt.Sprintf("[%s]: %s", sender.Name, message), sender, false)
}

func (cr *ChatRoom) SendPrivateMessage(message string, sender *User, receiverName string) {
	if receiver, ok := cr.users[receiverName]; ok {
		receiver.Receive(fmt.Sprintf("[ЛС от %s]: %s", sender.Name, message))
		sender.Receive(fmt.Sprintf("[Вы -> %s]: %s", receiver.Name, message))
	} else {
		sender.Receive(fmt.Sprintf("Пользователь '%s' не найден.", receiverName))
	}
}

func main() {
	chatRoom := NewChatRoom()

	user1 := NewUser("Алихан")
	user2 := NewUser("Влад")
	user3 := NewUser("Стас")
	user4 := NewUser("Андрей") 

	user1.JoinChat(chatRoom)
	user2.JoinChat(chatRoom)
	user3.JoinChat(chatRoom)

	// --- 2. Тестирование (Broadcast) ---
	fmt.Println("\n--- Тест 1: Общий чат ---")
	user1.Send("Всем привет!")
	replyMessage := fmt.Sprintf("Привет, %s!", user1.Name)
	user2.Send(replyMessage)

	fmt.Println("\n--- Тест 2: Личные сообщения (Расширение) ---")
	privateMsg := fmt.Sprintf("%s, ты здесь?", user2.Name)
	user3.SendPrivate(privateMsg, user2.Name)
	user1.SendPrivate("Его пока что нету", "Позже придет")

	fmt.Println("\n--- Тест 3: Выход из чата (Расширение) ---")
	user2.LeaveChat()
	user1.Send(fmt.Sprintf("Куда %s пропал?", user2.Name))

	fmt.Println("\n--- Тест 4: Ошибочная отправка без регистрации ---")
	user4.Send("Меня кто-нибудь слышит?")
}
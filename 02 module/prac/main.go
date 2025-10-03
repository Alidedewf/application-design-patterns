package main

import "fmt"

type User struct {
	Name  string
	Email string
	Role  string
}

type UserManager struct {
	users []User
}

func (um *UserManager) AddUser(user User) {
	um.users = append(um.users, user)
	fmt.Println("User added:", user.Name)
}

func (um *UserManager) RemoveUser(email string) {
	for i, u := range um.users {
		if u.Email == email {
			um.users = append(um.users[:i], um.users[i+1:]...)
			fmt.Println("User removed:", u.Name)
			return
		}
	}
	fmt.Println("User not found:", email)
}

func (um *UserManager) UpdateUser(email string, newName string, newRole string) {
	for i, u := range um.users {
		if u.Email == email {
			um.users[i].Name = newName
			um.users[i].Role = newRole
			fmt.Println("User updated:", newName)
			return
		}
	}
	fmt.Println("User not found:", email)
}

func main() {
	manager := UserManager{}

	user1 := User{Name: "Alihan", Email: "alihan@example.com", Role: "User"}
	user2 := User{Name: "Aruzhan", Email: "aruzhan@example.com", Role: "Admin"}

	manager.AddUser(user1)
	manager.AddUser(user2)

	manager.UpdateUser("alihan@example.com", "Alihan M.", "Admin")

	manager.RemoveUser("aruzhan@example.com")

	manager.RemoveUser("notfound@example.com") 
}
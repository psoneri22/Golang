package models

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Simulated database
var users = []User{
	{ID: 1, Username: "user1", Email: "user1@example.com"},
	{ID: 2, Username: "user2", Email: "user2@example.com"},
	{ID: 3, Username: "user3", Email: "user3@example.com"},
	{ID: 4, Username: "user4", Email: "user4@example.com"},
}

// GetAllUsers returns all users
func GetAllUsers() []User {
	return users
}
func CreateUser(user *User) error {
	return nil
}

package models

type Contact struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Simulated database
var contacts = []Contact{
	{ID: 1, Username: "contact1", Email: "contact1@example.com"},
	{ID: 2, Username: "contact2", Email: "contact2@example.com"},
	{ID: 3, Username: "contact3", Email: "contact3@example.com"},
	{ID: 4, Username: "contact4", Email: "contact4@example.com"},
}

// GetAllContacts returns all contacts
func GetAllContacts() []Contact {
	return contacts
}

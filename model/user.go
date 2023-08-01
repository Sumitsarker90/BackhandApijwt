package model

// Country struct for db table - country
type Users struct {
	Id       int    `json:"id"`
	Email    string `json:"email" `
	Password string `jspn:"password"`
}

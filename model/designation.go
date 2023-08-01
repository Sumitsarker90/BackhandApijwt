package model

// Country struct for db table - country
type Employee struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Designation string `json:"designation"`
	Email       string `json:"email" `
	Phone       string `json:"phone"`
	Address     string `json:"address"`
}

package models

type User struct {
	UserId   string `json:"user_id"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Role     Role   `json:"role"`
}

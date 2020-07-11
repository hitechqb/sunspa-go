package models

type Role struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (s Role) IsAdmin() bool {
	return s.Id == 999
}

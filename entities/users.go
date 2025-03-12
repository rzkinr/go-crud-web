package entities

type Users struct {
	Id       uint
	Username string `json:"username"`
	Password string `json:"password"`
}

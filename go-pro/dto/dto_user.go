package dto

type UserDto struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UsersDto []UserDto

package dto

type UserDto struct {
	ID       string `json:"id"`
	Username string `json:"username" validate:"required,min=4,max=100"`
	Email    string `json:"email" validate:"required,min=3,max=100"`
	Password string `json:"password" validate:"required,min=3,max=100"`
}

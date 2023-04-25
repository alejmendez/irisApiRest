package dto

type UserRequest struct {
	Username string `json:"username" validate:"required,min=4,max=100"`
	Email    string `json:"email" validate:"required,email,min=3,max=100"`
	Password string `json:"password" validate:"password"`
}

type UserUpdateRequest struct {
	Username string `json:"username" validate:"min=4,max=100"`
	Email    string `json:"email" validate:"email,min=3,max=100"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

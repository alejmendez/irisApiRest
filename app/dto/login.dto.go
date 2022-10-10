package dto

type LoginDto struct {
	Identity string `json:"identity" validate:"required"`
	Password string `json:"password" validate:"required"`
}

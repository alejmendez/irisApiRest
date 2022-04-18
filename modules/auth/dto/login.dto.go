package dto

type LoginDto struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}

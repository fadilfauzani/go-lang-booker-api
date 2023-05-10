package dtos

type RegisterDto struct {
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
	State    string `json:"state"`
	City     string `json:"city"`
	PinCode  int    `json:"pinCode"`
}
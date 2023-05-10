package dtos

type ProfileDto struct {
	ID         int    `json:"id"`
	Nama       string `json:"nama"`
	Email      string `json:"email"`
	Balance    int    `json:"balance"`
	State      string `json:"state"`
	City       string `json:"city"`
	PinCode    int    `json:"pinCode"`
}

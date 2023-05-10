package dtos

type GorDto struct {
	ID       int     `json:"id"`
	Nama     string  `json:"nama"`
	Alamat      string `json:"alamat"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	PhotoUrl string  `json:"photoUrl"`
	Rating   float32 `json:"rating"`
	Category string `json:"category"`

}
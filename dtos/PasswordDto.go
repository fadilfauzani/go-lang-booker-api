package dtos

type PasswordDto struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
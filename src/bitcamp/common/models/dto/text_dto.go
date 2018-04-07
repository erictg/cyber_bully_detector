package dto

type TextDTO struct{
	UserId int `json:"user_id"`
	Content string `json:"content"`
	Sent  bool `json:"sent"`
	OtherNumber string `json:"other_number"`
	Quit bool `json:"-"`
}
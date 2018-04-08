package models

type User struct{
	Id 			int 	`json:"id"`
	Name 		string 	`json:"name"`
	IsParent 	bool 	`json:"is_parent"`
	FCMId		string 	`json:"fcm_id"`
}

type FlaggedText struct{
	Id         int     `json:"id"`
	Content    string  `json:"content"`
	Confidence float32 `json:"score"`
	Sent       bool    `json:"sent"`
	UserId     int    `json:"user_id"`
	OtherNumber string `json:"other_number"`
}

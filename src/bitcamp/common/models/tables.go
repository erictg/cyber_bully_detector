package models

type User struct{
	Id 			int 	`json:"id"`
	Name 		string 	`json:"name"`
	IsParent 	bool 	`json:"is_parent"`
}

type FlaggedText struct{
	Id         int     `json:"id"`
	Content    string  `json:"content"`
	Confidence float32 `json:"score"`
	Sent       bool    `json:"sent"`
	UserId     bool    `json:"user_id"`
}

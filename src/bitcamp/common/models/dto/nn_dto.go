package dto

type NNReq struct {
	Content 	string 		`json:"content"`
}

//not_insult': float(prediction[0]), 'insult':
type NNResp struct{
	NotInsult	float32 `json:"not_insult"`
	Insult		float32 `json:"insult"`
}

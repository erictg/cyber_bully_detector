package dto

type UserCreateDTO struct {
	Name string `json:"name"`
	IsParent bool `json:"is_parent"`
}

type PairDTO struct{
	CId int `json:"c_id"`
	PId int `json:"p_id"`

}


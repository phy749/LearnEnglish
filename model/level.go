package model

type Level struct {
	Id           int     `grom:primaryKey" json:"level_id"`
	Level_number int     `json:"level_number"`
	Title        string  `json:"tile"`
	Description  *string `gorm:"type:text;default:null" json:"description,omitempty"`
	Skill        int     `json:"skill"`
}

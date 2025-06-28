package model

type Vocabulary struct {
	Id            int     `gorm:primaryKey" json : "id"`
	Word          string  `json:"word"`
	Meaning       *string `gorm:"type:text;default:null" json:"meaning,omitempty"`
	Example       *string `gorm:"type:text;default:null" json:"example,omitempty"`
	Pronunciation string  `json:"pronunciation"`
}

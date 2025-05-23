package model

type Role struct {
	Role_id     int     `gorm:primaryKey" json:"role_id"`
	Name        string  `json:"name"`
	Description *string `gorm:"type:text;default:null" json:"description,omitempty"`
}

package dataoject

import "time"

type UpdateImformationUser struct {
	Id        int        `json:"user_id"`
	Username  string     `json:"username"`
	Fullname  string     `json:"fullname"`
	Email     string     `json:"email"`
	Birthdate *time.Time `gorm:"type:date;default:CURRENT_DATE" json:"birthdate,omitempty"`
	Phone     string     `json:"phone"`
	Gender    *string    `gorm:"type:enum('male','female','other');default:other" json:"gender,omitempty"`
}

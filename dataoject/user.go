package dataoject

import "time"

type User struct {
	Username  string     `json:"username"`
	Fullname  string     `json:"fullname"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Birthdate *time.Time `gorm:"type:date;default:CURRENT_DATE" json:"birthdate,omitempty"`
	Phone     string     `json:"phone"`
	Gender    *string    `gorm:"type:enum('male','female','other');default:other" json:"gender,omitempty"`
}

package model

import "time"

type Useraccount struct {
	User_id    int        `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Username   string     `json:"username"`
	Fullname   string     `json:"fullname"`
	Email      string     `json:"email"`
	Password   string     `json:"password"`
	Birthdate  *time.Time `gorm:"type:date;default:CURRENT_DATE" json:"birthdate,omitempty"`
	Phone      string     `json:"phone"`
	Gender     *string    `gorm:"type:enum('male','female','other');default:other" json:"gender,omitempty"`
	Is_active  *string    `gorm:"type:enum('Y','N');default:'Y'" json:"is_active,omitempty"`
	Created_at *time.Time `gorm:"type:date;default:CURRENT_DATE" json:"create_at,omitempty"`

	RoleID int  `json:"role_id"` // Khóa ngoại
	Role   Role `gorm:"foreignKey:RoleID;references:Role_id" json:"role"`
}

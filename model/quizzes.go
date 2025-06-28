package model

type Quizzes struct {
	ID            int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	LessonID      int64  `gorm:"column:lesson_id" json:"lesson_id"`
	Question      string `gorm:"type:text" json:"question"`
	OptionA       string `gorm:"type:varchar(255);column:option_a" json:"option_a"`
	OptionB       string `gorm:"type:varchar(255);column:option_b" json:"option_b"`
	OptionC       string `gorm:"type:varchar(255);column:option_c" json:"option_c"`
	OptionD       string `gorm:"type:varchar(255);column:option_d" json:"option_d"`
	CorrectOption string `gorm:"type:char(1);column:correct_option" json:"correct_option"`
}

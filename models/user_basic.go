package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Identity         string `gorm:"column:identity;type varchar(36);" json:"identity"`
	Name             string `gorm:"column:name;type:varchar(100);" json:"name"`
	Passwrod         string `gorm:"column:passwrod;type:varchar(32);" json:"passwrod"`
	Phone            string `gorm:"column:phone;type:varchar(20);" json:"phone"`
	Mail             string `gorm:"column:mail;type:varchar(100);" json:"mail"`
	FinishProblemNum int    `gorm:"column:finish_problem_num;type:int(11);" json:"finish_problem_num"`
	SubmitNum        int    `gorm:"column:submit_num;type:int(11);" json:"submit_num"`
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

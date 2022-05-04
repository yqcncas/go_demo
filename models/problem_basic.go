package models

import (
	"gorm.io/gorm"
)

type ProblemBasic struct {
	gorm.Model
	Identity          string             `gorm:"column:identity;type varchar(36);" json:"identity"`
	Title             string             `gorm:"column:title;type: varchar(255);" json:"title"`
	Content           string             `gorm:"column:content;type:text;" json:"content"`
	MaxRuntime        int                `gorm:"column:max_runtime;type: int(11);" json:"max_runtime"`
	MaxMem            int                `gorm:"column:max_mem;type: int(11);" json:"max_mem"`
	ProblemCategories []*ProblemCategory `gorm:"foreignKey:problem_id;references:id"` // 关联问题分类表
}

func (table *ProblemBasic) TableName() string {
	return "problem_basic"
}

func GetProblemList(keyword string, categoryIdentity string) *gorm.DB {
	tx := DB.Model(new(ProblemBasic)).Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic").
		Where("title like ? OR Content like ?", "%"+keyword+"%", "%"+keyword+"%")
	if categoryIdentity != "" {
		tx.Joins("RIGHT JOIN problem_category pc on pc.problem_id = problem_basic.id").
			Where("pc.category_id = (SELECT cb.id FROM category_basic cb Where cb.identity = ?)", categoryIdentity)
	}

	return tx
}

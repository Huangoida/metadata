package model

import (
	"time"

	"gorm.io/gorm"
)

type DslInfoStruct struct {
	Id        int64          `gorm:"column:id"`
	Name      string         `gorm:"column:name"`
	Path      string         `gorm:"column:path"`
	Content   string         `gorm:"column:content"`
	Method    string         `gorm:"column:method"`
	CreatedAt time.Time      `gorm:"created_at;<-:create"`
	UpdatedAt time.Time      `gorm:"updated_at;<-:update"`
	Deleted   gorm.DeletedAt `gorm:"deleted"`
}

func (*DslInfoStruct) TableName() string {
	return "dsl_info"
}

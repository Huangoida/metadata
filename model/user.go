package model

import (
	"time"

	"gorm.io/gorm"
)

type UserStruct struct {
	Id        int64          `gorm:"column:id"`
	Name      string         `gorm:"column:name"`
	Pwd       string         `gorm:"column:pwd"`
	CreatedAt time.Time      `gorm:"created_at;<-:create"`
	UpdatedAt time.Time      `gorm:"updated_at;<-:update"`
	Deleted   gorm.DeletedAt `gorm:"deleted"`
}

func (*UserStruct) TableName() string {
	return "user"
}

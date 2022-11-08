package model

import (
	"gorm.io/gorm"
	"time"
)

type ServicesStruct struct {
	Id        int64          `gorm:"column:id"`
	Name      string         `gorm:"column:name"`
	Host      string         `gorm:"column:host"`
	Port      int            `gorm:"column:port"`
	Describes string         `gorm:"describes"`
	CreatedAt time.Time      `gorm:"created_at;<-:create"`
	UpdatedAt time.Time      `gorm:"updated_at;<-:update"`
	Deleted   gorm.DeletedAt `gorm:"deleted"`
}

func (*ServicesStruct) TableName() string {
	return "services"
}

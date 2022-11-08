package model

import (
	"gorm.io/gorm"
	"time"
)

type ApiStruct struct {
	Id             int64          `gorm:"column:id"`
	Name           string         `gorm:"column:name"`
	ServicesId     int64          `gorm:"column:services_id"`
	Path           string         `gorm:"column:path"`
	Protocol       string         `gorm:"column:protocol"`
	ConnectTimeout int            `gorm:"column:connect_timeout"`
	Retries        int            `gorm:"column:retries"`
	Status         string         `gorm:"column:status"`
	Tags           string         `gorm:"column:tags"`
	Method         string         `gorm:"column:method"`
	CreatedAt      time.Time      `gorm:"created_at;<-:create"`
	UpdatedAt      time.Time      `gorm:"updated_at;<-:update"`
	Deleted        gorm.DeletedAt `gorm:"deleted"`
}

func (*ApiStruct) TableName() string {
	return "api"
}

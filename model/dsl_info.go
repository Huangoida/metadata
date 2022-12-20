package model

import (
	"time"

	"gorm.io/gorm"
)

type DslInfoStruct struct {
	Id        int64          `gorm:"column:id" bson:"_id"`
	Name      string         `gorm:"column:name" bson:"name"`
	Path      string         `gorm:"column:path" bson:"path"`
	UserId    int64          `bson:"user_id"`
	Content   string         `gorm:"column:content" bson:"content"`
	Method    string         `gorm:"column:method" bson:"method"`
	CreatedAt time.Time      `gorm:"created_at;<-:create" bson:"created_at"`
	UpdatedAt time.Time      `gorm:"updated_at;<-:update" bson:"updated_at"`
	Deleted   gorm.DeletedAt `gorm:"deleted" bson:"deleted"`
}

func (*DslInfoStruct) TableName() string {
	return "dsl_info"
}

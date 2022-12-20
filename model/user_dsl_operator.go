package model

import (
	"gorm.io/gorm"
	"time"
)

type UserDslOperatorStruct struct {
	Id        int64          `gorm:"column:id"`
	UserId    int64          `gorm:"column:user_id"`
	Path      string         `gorm:"column:path"`
	DslId     int64          `gorm:"column:dsl_id"`
	Status    bool           `gorm:"column:status"`
	CreatedAt time.Time      `gorm:"created_at;<-:create"`
	UpdatedAt time.Time      `gorm:"updated_at;<-:update"`
	Deleted   gorm.DeletedAt `gorm:"deleted" bson:"deleted"`
}

func (*UserDslOperatorStruct) TableName() string {
	return "user_dsl_operator"
}

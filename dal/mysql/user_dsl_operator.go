package mysql

import (
	"context"
	"metadata/constant"
	"metadata/model"
)

func CreateUserDslOperator(ctx context.Context, operatorStruct model.UserDslOperatorStruct) error {
	err := GetDb().WithContext(ctx).Create(&operatorStruct).Error
	if err != nil {
		return err
	}
	return nil
}

func ListUserDslOperator(ctx context.Context, page, size int, id, userId, dslId int64, path string, status string, userOperator *[]model.UserDslOperatorStruct) (error, int64) {
	query := GetDb().WithContext(ctx).Table("user_dsl_operator")
	if id != 0 {
		query = query.Where("id = ?", id)
	}
	if userId != 0 {
		query = query.Where("user_id = ?", userId)
	}
	if dslId != 0 {
		query = query.Where("dsl_id = ?", dslId)
	}
	if path != "" {
		query = query.Where("path = ?", path)
	}

	if status != "" {
		if status == constant.BOOLEAN_FALSE {
			query = query.Where("status = false")
		} else if status == constant.BOOLEAN_TRUE {
			query = query.Where("status = true")
		}
	}

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return err, 0
	}

	if size != 0 && page != 0 {
		offset := (page - 1) * size
		query = query.Offset(offset).Limit(size)
	}

	if err := query.Debug().Find(&userOperator).Error; err != nil {

		return err, 0
	}
	return nil, count

}

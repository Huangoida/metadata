package mysql

import (
	"context"
	"metadata/model"
)

func CreateUser(ctx context.Context, user model.UserStruct) error {
	return GetDb().WithContext(ctx).Create(&user).Error
}

func ListUser(ctx context.Context, page, size int, name, id string, UserList *[]model.UserStruct) (error, int64) {
	query := GetDb().WithContext(ctx).Table("user")

	if name != "" {
		query = query.Where("name = ?", name)
	}
	// 这里不能提供密码查询，会造成严重的安全隐患
	if id != "" {
		query = query.Where("id = ?", id)
	}

	var count int64

	if size != 0 && page != 0 {
		offset := (page - 1) * size
		query = query.Offset(offset).Limit(size)
	}

	if err := query.Count(&count).Error; err != nil {
		return err, 0
	}

	if err := query.Find(&UserList).Error; err != nil {
		return err, 0
	}
	return nil, count
}

func UpdateUser(ctx context.Context, user model.UserStruct) error {
	err := GetDb().WithContext(ctx).Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(ctx context.Context, user model.UserStruct) error {
	err := GetDb().WithContext(ctx).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

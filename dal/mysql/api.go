package mysql

import (
	"context"
	"metadata/model"
)

func CreateApi(ctx context.Context, api model.ApiStruct) error {
	err := GetDb().WithContext(ctx).Create(&api).Error
	if err != nil {
		return err
	}
	return nil
}

func ListApi(ctx context.Context, page, size int, path, name, method, userId string, id, serviceId int64, apiList *[]model.ApiStruct) (error, int64) {
	query := GetDb().WithContext(ctx).Table("api")
	if path != "" {
		query = query.Where("path = ?", name)
	}
	if name != "" {
		query = query.Where("name = ?", name)
	}
	if method != "" {
		query = query.Where("method = ?", method)
	}

	if id != 0 {
		query = query.Where("id = ?", id)
	}
	if userId != "" {
		query = query.Where("user_id = ?", userId)
	}

	if serviceId != 0 {
		query = query.Where("services_id = ?", serviceId)
	}
	query.Where("deleted IS NULL")
	var count int64
	if err := query.Debug().Count(&count).Error; err != nil {
		return err, 0
	}

	if size != 0 && page != 0 {
		offset := (page - 1) * size
		query = query.Offset(offset).Limit(size)
	}

	if err := query.Find(&apiList).Error; err != nil {
		return err, 0
	}
	return nil, count
}

func UpdateApi(ctx context.Context, api model.ApiStruct) error {
	err := GetDb().WithContext(ctx).Save(&api).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteApi(ctx context.Context, api model.ApiStruct) error {
	err := GetDb().WithContext(ctx).Delete(&api).Error
	if err != nil {
		return err
	}
	return nil
}

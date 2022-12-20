package mysql

import (
	"context"
	"metadata/model"
)

func CreateServices(ctx context.Context, services model.ServicesStruct) error {
	return GetDb().WithContext(ctx).Create(&services).Error
}

func ListServices(ctx context.Context, page, size int, name, host, port, id, userId string, ServicesList *[]model.ServicesStruct) (error, int64) {
	query := GetDb().WithContext(ctx).Table("services")
	if name != "" {
		query = query.Where("name = ?", name)
	}
	if host != "" {
		query = query.Where("host = ?", host)
	}
	if port != "" {
		query = query.Where("port = ?", port)
	}
	if id != "" {
		query = query.Where("id = ?", id)
	}
	if userId != "" {
		query = query.Where("user_id = ?", userId)
	}
	var count int64
	if size != 0 && page != 0 {
		offset := (page - 1) * size
		query = query.Offset(offset).Limit(size)
	}
	if err := query.Count(&count).Error; err != nil {
		return err, 0
	}

	if err := query.Debug().Find(&ServicesList).Error; err != nil {

		return err, 0
	}
	return nil, count
}

func UpdateServices(ctx context.Context, services model.ServicesStruct) error {
	err := GetDb().WithContext(ctx).Save(&services).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteServices(ctx context.Context, services model.ServicesStruct) error {
	err := GetDb().WithContext(ctx).Delete(&services).Error
	if err != nil {
		return err
	}
	return nil
}

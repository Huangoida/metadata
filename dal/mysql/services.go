package mysql

import (
	"context"
	"gorm.io/gorm"
	"metadata/model"
)

func CreateServices(ctx context.Context, services model.ServicesStruct) error {
	return GetDb().WithContext(ctx).Create(&services).Error
}

func ListServices(ctx context.Context, page, size int, name, host, port, userId string, id []string, ServicesList *[]model.ServicesStruct) (error, int64) {
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
	if len(id) != 0 {
		if len(id) == 1 && id[0] != "" {
			query = query.Where("id IN ?", id)
		}
	}
	if userId != "" {
		query = query.Where("user_id = ?", userId)
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
	err := GetDb().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&services).Error; err != nil {
			return err
		}
		if err := tx.Where("services_id = ?", services.Id).Delete(&model.ApiStruct{}).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

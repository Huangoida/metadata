package mysql

import (
	"context"
	"gorm.io/gorm"
	"metadata/model"
)

func CreateParameters(ctx context.Context, parameters model.ParametersStruct) error {
	return GetDb().WithContext(ctx).Create(&parameters).Error
}

func CreateParametersBody(ctx context.Context, parametersBody []model.ParametersBodyStruct) error {
	return GetDb().WithContext(ctx).Create(&parametersBody).Error
}

func CreateParameterTransaction(ctx context.Context, parameters model.ParametersStruct, parametersBody []model.ParametersBodyStruct) error {
	err := GetDb().Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(&parameters).Error; err != nil {
			return err
		}

		if err := tx.WithContext(ctx).Create(&parametersBody).Error; err != nil {
			return err
		}

		return nil
	})
	return err
}

func ListParameter(ctx context.Context, page, size int, apiId, parameterId int64, userId string, parameterList *[]model.ParametersStruct) (error, int64) {
	query := GetDb().WithContext(ctx).Table("parameters")
	if apiId != 0 {
		query = query.Where("api_id = ?", apiId)
	}
	if parameterId != 0 {
		query = query.Where("id = ?", parameterId)
	}
	if userId != "" {
		query = query.Where("user_id = ?", userId)
	}

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return err, 0
	}

	if size != 0 && page != 0 {
		offset := (page - 1) * size
		query = query.Offset(offset).Limit(size)
	}

	if err := query.Find(&parameterList).Error; err != nil {
		return err, 0
	}
	return nil, count

}

func ListParameterBody(ctx context.Context, parameterId int64, parameterBodyList *[]model.ParametersBodyStruct) (error, int64) {
	query := GetDb().WithContext(ctx).Table("parameters_body")
	if parameterId != 0 {
		query = query.Where("id = ?", parameterId)
	}

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return err, 0
	}

	if err := query.Debug().Find(&parameterBodyList).Error; err != nil {
		return err, 0
	}
	return nil, count
}

func DeleteParameterBody(ctx context.Context, parameterId int64) error {
	return GetDb().WithContext(ctx).Table("parameters_body").Where("parameter_id = ?", parameterId).Delete(&model.ParametersBodyStruct{}).Error
}

func DeleteParameter(ctx context.Context, parameter model.ParametersStruct) error {
	return GetDb().WithContext(ctx).Delete(&parameter).Error
}

func UpdateParameter(ctx context.Context, parameter model.ParametersStruct) error {
	return GetDb().WithContext(ctx).Save(&parameter).Error
}

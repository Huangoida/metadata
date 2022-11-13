package dal

import (
	"context"
	"metadata/model"
)

func CreateDslInfo(ctx context.Context, dslInfo model.DslInfoStruct) error {
	err := GetDb().WithContext(ctx).Debug().Create(&dslInfo).Error
	if err != nil {
		return err
	}
	return nil
}

package mysql

import (
	"context"
	"metadata/model"
)

func CreateUserDslOperator(ctx context.Context, operatorStruct model.UserDslOperatorStruct) error {
	err := GetDb().WithContext(ctx).Create(&operatorStruct).Error
	if err != nil {
		return err
	}
	return nil
}

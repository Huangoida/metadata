package model

type ParametersBodyStruct struct {
	Id          int64  `gorm:"column:id"`
	ParameterId int64  `gorm:"column:parameter_id"`
	ParentId    int64  `gorm:"column:parent_id"`
	Key         string `gorm:"column:key"`
	Type        string `gorm:"column:type"`
}

func (*ParametersBodyStruct) TableName() string {
	return "parameters_body"
}

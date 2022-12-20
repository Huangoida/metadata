package model

type ParametersStruct struct {
	Id      int64  `gorm:"column:id"`
	ApiId   int64  `gorm:"column:api_id"`
	UserId  int64  `gorm:"column:user_id"`
	Key     string `gorm:"column:key"`
	Type    string `gorm:"column:type"`
	Value   string `gorm:"column:value"`
	Require bool   `gorm:"column:require"`
	Body    string `gorm:"column:body"`
}

func (*ParametersStruct) TableName() string {
	return "parameters"
}

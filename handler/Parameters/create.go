package Parameters

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"metadata/constant"
	"metadata/dal/mysql"
	"metadata/model"
	"metadata/util"
	"strconv"
)

type CreateParametersRequest struct {
	ApiId   int64  `json:"ApiId" binding:"required"`
	Type    string `json:"Type" binding:"required"`
	Key     string `json:"Key"`
	Value   string `json:"Value"`
	Require bool   `json:"Require"`
	Body    string `json:"Body"`
}

func Create(c *gin.Context) {
	userIdStr := c.GetHeader("UserId")
	var parameterRequest CreateParametersRequest
	if err := c.ShouldBindJSON(&parameterRequest); err != nil {
		logrus.Errorf("parameter invalid %v", err.Error())
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
		return
	}
	if parameterRequest.Type == "query" {
		queryTypeDealWith(c, parameterRequest)
		return
	}
	body := make(map[string]interface{})
	err := json.Unmarshal([]byte(parameterRequest.Body), &body)
	if err != nil {
		logrus.Errorf("parameter invalid %v", err.Error())
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
		return
	}

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		logrus.Errorf("parse userId failed %v", err.Error())
		util.ResponseError(c, 401, constant.USER_INVALID, "user invalid")
		return
	}

	parmeter := model.ParametersStruct{
		Id:      util.GenerateId(),
		ApiId:   parameterRequest.ApiId,
		Key:     parameterRequest.Key,
		Type:    parameterRequest.Type,
		UserId:  userId,
		Value:   parameterRequest.Value,
		Require: parameterRequest.Require,
		Body:    parameterRequest.Body,
	}

	var bodyList []model.ParametersBodyStruct
	err = ParseBody(0, parmeter.Id, body, &bodyList)
	if err != nil {
		logrus.Errorf("parameter invalid %v", err.Error())
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
		return
	}
	err = mysql.CreateParameterTransaction(c, parmeter, bodyList)
	if err != nil {
		logrus.Errorf("create ParametersBodyStruct failed %v", err.Error())
		util.ResponseError(c, 500, constant.CREATE_FAILED, "create ParametersBodyStruct failed")
		return
	}

	util.ResponseSuccess(c, "success")
}

func ParseBody(parentId int64, parmeterId int64, body map[string]interface{}, bodyList *[]model.ParametersBodyStruct) error {
	for key, value := range body {
		valueType := getValueType(value)
		bodyElement := model.ParametersBodyStruct{
			Id:          util.GenerateId(),
			ParameterId: parmeterId,
			ParentId:    parentId,
			Key:         key,
			Type:        valueType,
		}
		if valueType == "map" {
			v, ok := value.(map[string]interface{})
			if ok {
				err := ParseBody(bodyElement.Id, parmeterId, v, bodyList)
				if err != nil {
					return err
				}
			} else {
				return errors.New("cast map failed")
			}
		}
		*bodyList = append(*bodyList, bodyElement)
	}
	return nil
}

func getValueType(value interface{}) string {
	return fmt.Sprintf("%T", value)
}

func queryTypeDealWith(c *gin.Context, parameterRequest CreateParametersRequest) {
	if parameterRequest.Key == "" {
		logrus.Errorf("parameter invalid")
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
	}
	if parameterRequest.Value == "" {
		logrus.Errorf("parameter invalid")
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
	}
	parmeter := model.ParametersStruct{
		Id:      util.GenerateId(),
		ApiId:   parameterRequest.ApiId,
		Key:     parameterRequest.Key,
		Type:    parameterRequest.Type,
		Value:   parameterRequest.Value,
		Require: parameterRequest.Require,
		Body:    parameterRequest.Body,
	}

	err := mysql.CreateParameters(c, parmeter)
	if err != nil {
		logrus.Errorf("create Parameter failed %v", err.Error())
		util.ResponseError(c, 500, constant.CREATE_FAILED, "create Parameter failed")
		return
	}

	util.ResponseSuccess(c, "success")

}

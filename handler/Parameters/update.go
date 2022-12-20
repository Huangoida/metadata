package Parameters

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"metadata/constant"
	"metadata/dal/mysql"
	"metadata/model"
	"metadata/util"
)

type UpdateParameterRequest struct {
	ApiId       int64  `json:"ApiId" binding:"required"`
	ParameterId int64  `json:"ParameterId" binding:"required"`
	Type        string `json:"Type"`
	Key         string `json:"Key"`
	Value       string `json:"Value"`
	IsRequire   bool   `json:"IsRequire"`
	Require     bool   `json:"Require"`
	Body        string `json:"Body"`
}

func Update(c *gin.Context) {
	userId := c.GetHeader("UserId")
	var updateParameterRequest UpdateParameterRequest
	if err := c.ShouldBindQuery(&updateParameterRequest); err != nil {
		logrus.Errorf("parameter invalid %v", err.Error())
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
		return
	}
	var parameterList []model.ParametersStruct
	err, count := mysql.ListParameter(c, 0, 0, updateParameterRequest.ApiId, updateParameterRequest.ParameterId, userId, &parameterList)
	if err != nil {
		logrus.Errorf("parameter invalid %v", err.Error())
		util.ResponseError(c, 500, constant.SEARCH_FAILED, "search failed")
		return
	}
	if count != 1 {
		logrus.Errorf("parameter not found")
		util.ResponseError(c, 400, constant.SEARCH_NOT_FOUND, "parameter not found")
		return
	}
	parameter := parameterList[0]
	if updateParameterRequest.IsRequire {
		parameter.Require = updateParameterRequest.Require
	}
	if updateParameterRequest.Type != "" {
		if updateParameterRequest.Type == "body" {
			if updateParameterRequest.Body == "" {
				logrus.Errorf("parameter invalid %v", err.Error())
				util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
				return
			}
		} else {
			if updateParameterRequest.Key == "" {
				logrus.Errorf("parameter invalid %v", err.Error())
				util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
				return
			}
			if updateParameterRequest.Value == "" {
				logrus.Errorf("parameter invalid %v", err.Error())
				util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
				return
			}
		}
		parameter.Type = updateParameterRequest.Type
	}

	if updateParameterRequest.Key != "" && parameter.Type == "query" {
		parameter.Key = updateParameterRequest.Key
	}
	if updateParameterRequest.Value != "" && parameter.Type == "query" {
		parameter.Value = updateParameterRequest.Value
	}

	if updateParameterRequest.Body != "" && parameter.Body == "body" {
		parameter.Body = updateParameterRequest.Body
		body := make(map[string]interface{})
		err = json.Unmarshal([]byte(updateParameterRequest.Body), &body)
		if err != nil {
			logrus.Errorf("parameter invalid %v", err.Error())
			util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
			return
		}
		var bodyList []model.ParametersBodyStruct
		err = ParseBody(0, parameter.Id, body, &bodyList)
		if err != nil {
			logrus.Errorf("parameter invalid %v", err.Error())
			util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
			return
		}
		err := mysql.DeleteParameterBody(c, parameter.Id)
		if err != nil {
			logrus.Errorf("delete failed %v", err.Error())
			util.ResponseError(c, 500, constant.DELETE_FAILED, "delete failed")
			return
		}
		err = mysql.CreateParametersBody(c, bodyList)
		if err != nil {
			logrus.Errorf("create ParametersBody failed %v", err.Error())
			util.ResponseError(c, 500, constant.CREATE_FAILED, "create ParametersBody failed")
			return
		}
	}

	err = mysql.UpdateParameter(c, parameter)
	if err != nil {
		logrus.Errorf("update api failed %v", err.Error())
		util.ResponseError(c, 500, constant.UPDATE_FAILED, "update api failed")
		return
	}

	util.ResponseSuccess(c, "success")
}

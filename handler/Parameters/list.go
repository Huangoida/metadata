package Parameters

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"metadata/constant"
	"metadata/dal/mysql"
	"metadata/model"
	"metadata/util"
	"strconv"
)

func List(c *gin.Context) {
	page, size := util.ValidateOffsetAndPage(c)
	apiIdStr := c.Query("ApiId")
	if apiIdStr == "" {
		logrus.Errorf("ApiId invalid ")
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "ApiId invalid")
		return
	}
	apiId, err := strconv.ParseInt(apiIdStr, 10, 64)
	if err != nil {
		logrus.Errorf("ApiId invalid ")
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "ApiId invalid")
		return
	}
	parameterIdStr := c.Query("parameterId")
	if parameterIdStr == "" {
		parameterIdStr = "0"
	}
	parameterId, err := strconv.ParseInt(parameterIdStr, 10, 64)
	if err != nil {
		logrus.Errorf("parameter invalid %v", err.Error())
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
		return
	}
	var apiList []model.ApiStruct
	err, count := mysql.ListApi(c, page, size, "", "", "", apiId, 0, &apiList)
	if err != nil {
		logrus.Errorf("parameter invalid %v", err.Error())
		util.ResponseError(c, 500, constant.SEARCH_FAILED, "search failed")
		return
	}
	if count == 0 {
		logrus.Errorf("api not found")
		util.ResponseError(c, 400, constant.SEARCH_NOT_FOUND, "api not found")
	}
	var parameterList []model.ParametersStruct
	err, count = mysql.ListParameter(c, page, size, apiId, parameterId, &parameterList)
	if err != nil {
		logrus.Errorf("parameter invalid %v", err.Error())
		util.ResponseError(c, 500, constant.SEARCH_FAILED, "search failed")
		return
	}
	if count == 0 {
		util.ResponseSuccess(c, map[string]interface{}{
			"count": count,
			"res":   nil,
		})
		return
	}
	var res []map[string]interface{}
	for _, parameter := range parameterList {
		if parameter.Type == "body" {
			var parameterBodyList []model.ParametersBodyStruct
			err, _ = mysql.ListParameterBody(c, parameterId, &parameterBodyList)
			if err != nil {
				logrus.Errorf("parameter invalid %v", err.Error())
				util.ResponseError(c, 500, constant.SEARCH_FAILED, "search failed")
				return
			}
			body := generateBody(0, parameterBodyList)
			res = append(res, map[string]interface{}{
				"Type":     parameter.Type,
				"Require":  parameter.Require,
				"BodyType": body,
				"Boyd":     parameter.Body,
			})
		} else {
			res = append(res, map[string]interface{}{
				"Type":    parameter.Type,
				"Require": parameter.Require,
				"Key":     parameter.Key,
				"Value":   parameter.Value,
			})
		}
	}
	util.ResponseSuccess(c, map[string]interface{}{
		"count": count,
		"res":   res,
	})
}

func generateBody(parentId int64, parameterBodyList []model.ParametersBodyStruct) map[string]interface{} {
	body := make(map[string]interface{})
	for _, parametersBody := range parameterBodyList {
		if parametersBody.ParentId == parentId {
			if parametersBody.Type == "map" {
				body[parametersBody.Key] = generateBody(parametersBody.Id, parameterBodyList)
			} else {
				body[parametersBody.Key] = parametersBody.Type
			}
		}
	}
	return body
}

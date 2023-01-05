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

func Delete(c *gin.Context) {
	userId := c.GetHeader("UserId")
	apiIdstr := c.Query("ApiId")
	if apiIdstr == "" {
		logrus.Errorf("parameter invalid")
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
		return
	}
	parameterIdstr := c.Query("ParameterId")
	if parameterIdstr == "" {
		logrus.Errorf("parameter invalid")
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
		return
	}

	var apiId, parameterId int64
	if apiIdstr == "" {
		apiIdstr = "0"
	}
	apiId, err := strconv.ParseInt(apiIdstr, 10, 64)
	if err != nil {
		logrus.Errorf("parameter invalid %v", err.Error())
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
		return
	}
	if parameterIdstr == "" {
		parameterIdstr = "0"
	}
	parameterId, err = strconv.ParseInt(parameterIdstr, 10, 64)
	if err != nil {
		logrus.Errorf("parameter invalid %v", err.Error())
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
		return
	}
	var parameterList []model.ParametersStruct
	err, count := mysql.ListParameter(c, 0, 0, apiId, parameterId, userId, &parameterList)
	if err != nil {
		logrus.Errorf("parameter invalid %v", err.Error())
		util.ResponseError(c, 500, constant.SEARCH_FAILED, "search failed")
		return
	}

	if count == 0 {
		util.ResponseSuccess(c, "success")
		return
	}

	parameter := parameterList[0]
	if parameter.Type == "body" {
		err := mysql.DeleteParameterBody(c, parameter.Id)
		if err != nil {
			logrus.Errorf("delete failed %v", err.Error())
			util.ResponseError(c, 500, constant.DELETE_FAILED, "delete failed")
			return
		}
	}
	err = mysql.DeleteParameter(c, parameter)
	if err != nil {
		logrus.Errorf("delete failed %v", err.Error())
		util.ResponseError(c, 500, constant.DELETE_FAILED, "delete failed")
		return
	}

	util.ResponseSuccess(c, "success")
}

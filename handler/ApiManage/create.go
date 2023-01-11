package ApiManage

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"metadata/constant"
	"metadata/dal/mysql"
	"metadata/model"
	"metadata/util"
	"strconv"
)

type APIRequestStruct struct {
	Path           string `json:"Path" binding:"required"`
	ServiceId      string `json:"ServiceId" binding:"required"`
	Protocol       string `json:"Protocol" binding:"required"`
	Name           string `json:"Name" binding:"required"`
	Tags           string `json:"Tags"`
	Method         string `json:"Method" binding:"required"`
	ConnectTimeout int    `json:"ConnectTimeout" `
	Retries        int    `json:"Retries" `
}

func Create(c *gin.Context) {
	userIdStr := c.GetHeader("UserId")
	var apiRequest APIRequestStruct
	if err := c.ShouldBindJSON(&apiRequest); err != nil {
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
	serviceIdInt64, err := strconv.ParseInt(apiRequest.ServiceId, 10, 64)
	if err != nil {
		logrus.Errorf("parse ServiceId failed %v", err.Error())
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
		return
	}
	var services []model.ServicesStruct
	err, count := mysql.ListServices(c, 0, 0, "", "", "", userIdStr, []string{apiRequest.ServiceId}, &services)
	if err != nil {
		logrus.Errorf("search services failed %v", err.Error())
		util.ResponseError(c, 500, constant.CREATE_FAILED, "search services failed")
		return
	}
	if count == 0 {
		logrus.Errorf("parse ServiceId failed %v", err.Error())
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "ServiceId not found")
		return
	}

	api := model.ApiStruct{
		Id:             util.GenerateId(),
		Name:           apiRequest.Name,
		ServicesId:     serviceIdInt64,
		Path:           apiRequest.Path,
		Protocol:       apiRequest.Protocol,
		Status:         "disabled",
		UserId:         userId,
		Tags:           apiRequest.Tags,
		Method:         apiRequest.Method,
		ConnectTimeout: apiRequest.ConnectTimeout,
		Retries:        apiRequest.Retries,
	}

	err = mysql.CreateApi(c, api)
	if err != nil {
		logrus.Errorf("create services failed %v", err.Error())
		util.ResponseError(c, 500, constant.CREATE_FAILED, "create services failed")
		return
	}

	util.ResponseSuccess(c, map[string]interface{}{
		"ApiId": api.Id,
	})
}

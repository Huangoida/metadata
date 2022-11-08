package ServicesManage

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"metadata/constant"
	"metadata/dal"
	"metadata/model"
	"metadata/util"
)

type ServicesRequestStruct struct {
	Name     string `json:"Name" binding:"required"`
	Host     string `json:"Host" binding:"required"`
	Port     int    `json:"Port" binding:"gte=1,lte=65535"`
	Describe string `json:"Describe"`
}

func Create(c *gin.Context) {

	var servicesRequest ServicesRequestStruct
	if err := c.ShouldBindJSON(&servicesRequest); err != nil {
		logrus.Errorf("parameter invalid %v", err.Error())
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
		return
	}

	services := model.ServicesStruct{
		Id:        util.GenerateId(),
		Name:      servicesRequest.Name,
		Host:      servicesRequest.Host,
		Port:      servicesRequest.Port,
		Describes: servicesRequest.Describe,
	}

	err := dal.CreateServices(c, services)
	if err != nil {
		logrus.Errorf("create services failed %v", err.Error())
		util.ResponseError(c, 500, constant.CREATE_FAILED, "create services failed")
		return
	}

	util.ResponseSuccess(c, "success")
}

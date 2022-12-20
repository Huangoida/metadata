package ServicesManage

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"metadata/constant"
	"metadata/dal/mysql"
	"metadata/model"
	"metadata/util"
	"strconv"
)

type ServicesUpdateResponse struct {
	Id       int64  `form:"Id" binding:"required"`
	Name     string `form:"Name"`
	Host     string `form:"Host"`
	Port     int    `form:"Port"`
	Describe string `form:"Describe"`
}

func Update(c *gin.Context) {
	userId := c.GetHeader("UserId")
	var servicesResponse ServicesUpdateResponse
	if err := c.ShouldBindQuery(&servicesResponse); err != nil {
		logrus.Errorf("parameter invalid %v", err.Error())
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
		return
	}
	var servicesList []model.ServicesStruct
	err, count := mysql.ListServices(c, 0, 0, "", "", "", strconv.FormatInt(servicesResponse.Id, 10), userId, &servicesList)
	if err != nil {
		logrus.Errorf("search failed %v", err.Error())
		util.ResponseError(c, 500, constant.SEARCH_FAILED, "search failed")
		return
	}
	if count == 0 {
		util.ResponseError(c, 401, constant.SEARCH_NOT_FOUND, "search not found")
		return
	}
	services := servicesList[0]
	if servicesResponse.Port != 0 {
		services.Port = servicesResponse.Port
	}
	if servicesResponse.Name != "" {
		services.Name = servicesResponse.Name
	}
	if servicesResponse.Host != "" {
		services.Host = servicesResponse.Host
	}
	if servicesResponse.Describe != "" {
		services.Describes = servicesResponse.Describe
	}
	err = mysql.UpdateServices(c, services)
	if err != nil {
		logrus.Errorf("update service failed %v", err.Error())
		util.ResponseError(c, 500, constant.UPDATE_FAILED, "update service failed")
		return
	}

	util.ResponseSuccess(c, "success")
}

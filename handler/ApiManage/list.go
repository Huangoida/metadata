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

type apiListResStruct struct {
	model.ApiStruct
	ServiceName string
}

func List(c *gin.Context) {
	userId := c.GetHeader("UserId")
	page, size := util.ValidateOffsetAndPage(c)
	path := c.Query("Path")
	serviceIdStr := c.Query("ServiceId")
	if serviceIdStr == "" {
		serviceIdStr = "0"
	}
	serviceId, err := strconv.ParseInt(serviceIdStr, 10, 64)
	if err != nil {
		logrus.Errorf("parameter invalid %v", err.Error())
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
		return
	}

	var id int64
	idStr := c.Query("id")
	if idStr == "" {
		idStr = "0"
	}
	id, err = strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		logrus.Errorf("parameter invalid %v", err.Error())
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
		return
	}

	name := c.Query("Name")
	method := c.Query("Method")
	var apiList []model.ApiStruct

	err, count := mysql.ListApi(c, page, size, path, name, method, userId, id, serviceId, &apiList)
	if err != nil {
		logrus.Errorf("search failed %v", err.Error())
		util.ResponseError(c, 500, constant.SEARCH_FAILED, "search failed")
		return
	}

	if count == 0 {
		util.ResponseSuccess(c, map[string]interface{}{
			"count": count,
			"res":   apiList,
		})
		return
	}
	var servicesIdList []string
	for _, apiStruct := range apiList {
		servicesIdList = append(servicesIdList, strconv.FormatInt(apiStruct.ServicesId, 10))
	}

	var servicesList []model.ServicesStruct
	err, _ = mysql.ListServices(c, 0, 0, "", "", "", userId, servicesIdList, &servicesList)
	if err != nil {
		logrus.Errorf("search failed %v", err.Error())
		util.ResponseError(c, 500, constant.SEARCH_FAILED, "search failed")
		return
	}
	serviceNameMap := make(map[int64]string)
	for _, servicesStruct := range servicesList {
		serviceNameMap[servicesStruct.Id] = servicesStruct.Name
	}

	var resList []apiListResStruct
	for _, apiStruct := range apiList {
		sereviceName := serviceNameMap[apiStruct.ServicesId]
		if sereviceName != "" {
			resList = append(resList, apiListResStruct{
				ApiStruct:   apiStruct,
				ServiceName: sereviceName,
			})
		}
	}
	util.ResponseSuccess(c, map[string]interface{}{
		"count": count,
		"res":   resList,
	})
}

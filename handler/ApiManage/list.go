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

func List(c *gin.Context) {
	userId := c.GetHeader("UserId")
	page, size := util.ValidateOffsetAndPage(c)
	path := c.Query("Path")
	serviceIdStr := c.Query("ServiceId")
	if serviceIdStr == "" {
		logrus.Errorf("parameter invalid ")
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
		return
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
		logrus.Errorf("parameter invalid %v", err.Error())
		util.ResponseError(c, 500, constant.SEARCH_FAILED, "search failed")
		return
	}

	util.ResponseSuccess(c, map[string]interface{}{
		"count": count,
		"res":   apiList,
	})
}

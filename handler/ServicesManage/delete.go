package ServicesManage

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"metadata/constant"
	"metadata/dal/mysql"
	"metadata/model"
	"metadata/util"
)

func Delete(c *gin.Context) {
	userId := c.GetHeader("UserId")
	idstr := c.Query("Id")
	if idstr == "" {
		logrus.Errorf("parameter invalid")
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
		return
	}

	var servicesList []model.ServicesStruct
	err, count := mysql.ListServices(c, 0, 0, "", "", "", userId, []string{idstr}, &servicesList)
	if err != nil {
		logrus.Errorf("search service failed %v", err.Error())
		util.ResponseError(c, 500, constant.SEARCH_FAILED, "search failed")
		return
	}
	if count == 0 {
		util.ResponseError(c, 401, constant.SEARCH_NOT_FOUND, "search not found")
		return
	}
	services := servicesList[0]
	var apiList []model.ApiStruct
	err, count = mysql.ListApi(c, 0, 0, "", "", "", userId, 0, services.Id, &apiList)
	if err != nil {
		logrus.Errorf("search api failed %v", err.Error())
		util.ResponseError(c, 500, constant.DELETE_FAILED, "delete failed")
		return
	}

	err = mysql.DeleteServices(c, services)
	if err != nil {
		logrus.Errorf("delete failed %v", err.Error())
		util.ResponseError(c, 500, constant.DELETE_FAILED, "delete failed")
		return
	}

	util.ResponseSuccess(c, "success")
}

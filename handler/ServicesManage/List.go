package ServicesManage

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"metadata/constant"
	"metadata/dal/mysql"
	"metadata/model"
	"metadata/util"
)

func List(c *gin.Context) {
	userId := c.GetHeader("UserId")
	page, size := util.ValidateOffsetAndPage(c)
	name := c.Query("Name")
	hostName := c.Query("HostName")
	port := c.Query("Port")
	id := c.Query("Id")
	var servicesList []model.ServicesStruct
	err, count := mysql.ListServices(c, page, size, name, hostName, port, id, userId, &servicesList)
	if err != nil {
		logrus.Errorf("parameter invalid %v", err.Error())
		util.ResponseError(c, 500, constant.SEARCH_FAILED, "search failed")
		return
	}

	util.ResponseSuccess(c, map[string]interface{}{
		"count": count,
		"res":   servicesList,
	})

}

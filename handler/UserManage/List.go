package UserManage

import (
	"metadata/constant"
	"metadata/dal/mysql"
	"metadata/model"
	"metadata/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func List(c *gin.Context) {

	page, size := util.ValidateOffsetAndPage(c)
	name := c.Query("Name")
	id := c.Query("Id")
	var userList []model.UserStruct
	err, count := mysql.ListUser(c, page, size, name, id, &userList)
	if err != nil {
		logrus.Errorf("parameter invalid %v", err.Error())
		util.ResponseError(c, 500, constant.SEARCH_FAILED, "search failed")
		return
	}

	util.ResponseSuccess(c, map[string]interface{}{
		"count": count,
		"res":   userList,
	})

}

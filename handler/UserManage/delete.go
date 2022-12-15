package UserManage

import (
	"metadata/constant"
	"metadata/dal/mysql"
	"metadata/model"
	"metadata/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Delete(c *gin.Context) {

	idstr := c.Query("Id")
	if idstr == "" {
		logrus.Errorf("parameter invalid")
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
		return
	}

	var userList []model.UserStruct
	err, count := mysql.ListUser(c, 0, 0, "", idstr, &userList)
	if err != nil {
		logrus.Errorf("search failed %v", err.Error())
		util.ResponseError(c, 500, constant.SEARCH_FAILED, "search failed")
		return
	}
	if count == 0 {
		util.ResponseError(c, 401, constant.SEARCH_NOT_FOUND, "search not found")
		return
	}
	user := userList[0]
	err = mysql.DeleteUser(c, user)
	if err != nil {
		logrus.Errorf("delete failed %v", err.Error())
		util.ResponseError(c, 500, constant.DELETE_FAILED, "delete failed")
		return
	}

	util.ResponseSuccess(c, "success")
}

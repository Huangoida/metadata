package UserManage

import (
	"metadata/constant"
	"metadata/dal/mysql"
	"metadata/model"
	"metadata/util"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserUpdateResponse struct {
	Id   int64  `form:"Id" binding:"required"`
	Name string `form:"Name"`
	Pwd  string `form:"Pwd"`
}

func Update(c *gin.Context) {

	var userResponse UserUpdateResponse
	if err := c.ShouldBindQuery(&userResponse); err != nil {
		logrus.Errorf("parameter invalid %v", err.Error())
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
		return
	}
	var userList []model.UserStruct
	err, count := mysql.ListUser(c, 0, 0, "", strconv.FormatInt(userResponse.Id, 10), &userList)
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
	if userResponse.Name != "" {
		user.Name = userResponse.Name
	}
	encrypted, _ := util.GetPwd(userResponse.Pwd)
	if userResponse.Pwd != "" {
		user.Pwd = string(encrypted)
	}
	err = mysql.UpdateUser(c, user)
	if err != nil {
		logrus.Errorf("update user failed %v", err.Error())
		util.ResponseError(c, 500, constant.UPDATE_FAILED, "update user failed")
		return
	}

	util.ResponseSuccess(c, "success")
}

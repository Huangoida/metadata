package common

import (
	"metadata/constant"
	"metadata/dal/mysql"
	"metadata/handler/UserManage"
	"metadata/model"
	"metadata/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Login(c *gin.Context) {

	var userRequest UserManage.UserRequestStruct
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logrus.Errorf("parameter invalid %v", err.Error())
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
		return
	}

	// 查找用户是否存在
	var userList []model.UserStruct
	err, count := mysql.ListUser(c, 0, 1, userRequest.Name, "", &userList)
	if err != nil {
		logrus.Errorf("login failed, search failed %v", err.Error())
		util.ResponseError(c, 500, constant.SEARCH_FAILED, "login failed, search failed")
		return
	}
	if count == 0 {
		logrus.Errorf("login failed, User doesn't exist %v", err.Error())
		util.ResponseError(c, 500, constant.SEARCH_NOT_FOUND, "login failed, User doesn't exist")
		return
	}
	// 比较密码
	if util.ComparePwd(userList[0].Pwd, userRequest.Pwd) {
		logrus.Debugf("%v login", userList[0].Id)
		// 密码正确， 生成token
		token, _ := util.GenToken(userList[0].Id)
		util.ResponseSuccess(c, map[string]interface{}{
			"token": token,
		})
		return
	} else {
		logrus.Errorf("Password is not correct %v", err.Error())
		util.ResponseError(c, 500, constant.CREATE_FAILED, "Password is not correct")
		return
	}

}

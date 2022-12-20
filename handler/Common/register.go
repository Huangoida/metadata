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

func Register(c *gin.Context) {

	var userRequest UserManage.UserRequestStruct
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logrus.Errorf("parameter invalid %v", err.Error())
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
		return
	}

	// encrypted : 已加密的密码
	encrypted, _ := util.GetPwd(userRequest.Pwd)
	user := model.UserStruct{
		Id:   util.GenerateId(),
		Name: userRequest.Name,
		Pwd:  string(encrypted),
	}

	// 查找用户是否存在
	var userList []model.UserStruct
	err, count := mysql.ListUser(c, 0, 1, userRequest.Name, "", &userList)
	if err != nil {
		logrus.Errorf("register failed, search failed %v", err.Error())
		util.ResponseError(c, 500, constant.SEARCH_FAILED, "register failed, search failed")
		return
	}
	if count >= 1 {
		logrus.Errorf("create user failed, user already exists  %v", err.Error())
		util.ResponseError(c, 500, constant.CREATE_FAILED, "create user failed, User already exists")
		return
	}

	err = mysql.CreateUser(c, user)
	if err != nil {
		logrus.Errorf("create user failed %v", err.Error())
		util.ResponseError(c, 500, constant.CREATE_FAILED, "create user failed")
		return
	}

	util.ResponseSuccess(c, "success")
}

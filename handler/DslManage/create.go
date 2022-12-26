package DslManage

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"metadata/constant"
	"metadata/dal/mongo"
	"metadata/dal/mysql"
	"metadata/model"
	"metadata/util"
	"strconv"
)

type DSLRequestStruct struct {
	Path    string `json:"Path" binding:"required"`
	Method  string `json:"Method" binding:"required"`
	Content string `json:"Content" binding:"required"`
	Name    string `json:"Name" binding:"required"`
}

func Create(c *gin.Context) {
	userIdStr := c.GetHeader("UserId")

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		logrus.Errorf("parse userId failed %v", err.Error())
		util.ResponseError(c, 401, constant.USER_INVALID, "user invalid")
		return
	}

	var dslRequest DSLRequestStruct
	if err := c.ShouldBindJSON(&dslRequest); err != nil {
		logrus.Errorf("Dsl info invalid %v", err.Error())
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
		return
	}
	dsl := model.DslInfoStruct{
		Id:      util.GenerateId(),
		UserId:  userId,
		Name:    dslRequest.Name,
		Path:    dslRequest.Path,
		Content: dslRequest.Content,
		Method:  dslRequest.Method,
	}

	userOperate := model.UserDslOperatorStruct{
		Id:     util.GenerateId(),
		UserId: userId,
		Path:   dslRequest.Path,
		Name:   dslRequest.Name,
		Method: dslRequest.Method,
		DslId:  dsl.Id,
		Status: true,
	}
	err = mysql.CreateUserDslOperator(c, userOperate)
	if err != nil {
		logrus.Errorf("create user operator failed %v", err.Error())
		util.ResponseError(c, 500, constant.CREATE_FAILED, "create user operator failed")
		return
	}

	dsl.APIId = userOperate.Id

	err = mongo.CreateDslInfo(c, dsl)
	if err != nil {
		logrus.Errorf("create Dsl info failed %v", err.Error())
		util.ResponseError(c, 500, constant.CREATE_FAILED, "create Dsl info failed")
		return
	}

	util.ResponseSuccess(c, "success")
}

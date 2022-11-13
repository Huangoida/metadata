package Dslmanage

import (
	"metadata/constant"
	"metadata/dal"
	"metadata/model"
	"metadata/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type DSLRequestStruct struct {
	Path    string `json:"Path" binding:"required"`
	Method  string `json:"Method" binding:"required"`
	Content string `json:"Content" binding:"required"`
	Name    string `json:"Name" binding:"required"`
}

func Create(c *gin.Context) {
	var dslRequest DSLRequestStruct
	if err := c.ShouldBindJSON(&dslRequest); err != nil {
		logrus.Errorf("Dsl info invalid %v", err.Error())
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
		return
	}
	dsl := model.DslInfoStruct{
		Id:      util.GenerateId(),
		Name:    dslRequest.Name,
		Path:    dslRequest.Path,
		Content: dslRequest.Content,
		Method:  dslRequest.Method,
	}

	err := dal.CreateDslInfo(c, dsl)
	if err != nil {
		logrus.Errorf("create Dsl info failed %v", err.Error())
		util.ResponseError(c, 500, constant.CREATE_FAILED, "create services failed")
		return
	}

	util.ResponseSuccess(c, "success")
}

package ApiManage

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"metadata/constant"
	"metadata/dal"
	"metadata/model"
	"metadata/util"
	"strconv"
)

func Delete(c *gin.Context) {

	idstr := c.Query("Id")
	if idstr == "" {
		logrus.Errorf("parameter invalid")
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
		return
	}
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		logrus.Errorf("parameter invalid %v", err.Error())
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "parameter invalid")
		return
	}
	var apiList []model.ApiStruct
	err, count := dal.ListApi(c, 0, 0, "", "", "", id, 0, &apiList)
	if err != nil {
		logrus.Errorf("search failed %v", err.Error())
		util.ResponseError(c, 500, constant.SEARCH_FAILED, "search failed")
		return
	}
	if count == 0 {
		util.ResponseError(c, 401, constant.SEARCH_NOT_FOUND, "search not found")
		return
	}
	api := apiList[0]
	err = dal.DeleteApi(c, api)
	if err != nil {
		logrus.Errorf("delete failed %v", err.Error())
		util.ResponseError(c, 500, constant.DELETE_FAILED, "delete failed")
		return
	}

	util.ResponseSuccess(c, "success")
}

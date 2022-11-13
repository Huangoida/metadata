package Dslmanage

import (
	"metadata/constant"
	"metadata/dal"
	"metadata/model"
	"metadata/util"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func List(c *gin.Context) {

	page, size := util.ValidateOffsetAndPage(c)
	path := c.Query("Path")

	var id int64
	idStr := c.Query("id")
	if idStr == "" {
		idStr = "0"
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		logrus.Errorf("dsl info invalid %v", err.Error())
		util.ResponseError(c, 401, constant.PARAMETER_INVALID, "dsl info invalid")
		return
	}

	name := c.Query("Name")
	method := c.Query("Method")
	var dslList []model.DslInfoStruct

	err, count := dal.ListDslInfo(c, page, size, path, name, method, id, serviceId, &apiList)
	if err != nil {
		logrus.Errorf("parameter invalid %v", err.Error())
		util.ResponseError(c, 500, constant.SEARCH_FAILED, "search failed")
		return
	}

	util.ResponseSuccess(c, map[string]interface{}{
		"count": count,
		"res":   apiList,
	})
}

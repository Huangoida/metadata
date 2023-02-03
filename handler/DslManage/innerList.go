package DslManage

import (
	"metadata/constant"
	"metadata/dal/mongo"
	"metadata/dal/mysql"
	"metadata/model"
	"metadata/util"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func InnerList(c *gin.Context) {

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
	content := c.Query("Content")

	var userDsl []model.UserDslOperatorStruct
	err, i := mysql.ListUserDslOperator(c, 0, 0, id, 0, 0, path, constant.BOOLEAN_TRUE, &userDsl)
	if err != nil {
		logrus.Errorf("parameter invalid %v", err.Error())
		util.ResponseError(c, 500, constant.SEARCH_FAILED, "search failed")
		return
	}
	if i == 0 {
		logrus.Debugf("user operator not found")
		util.ResponseSuccess(c, map[string]interface{}{
			"count": i,
			"res":   nil,
		})
		return
	}

	var dslList []model.DslInfoStruct
	err, count := mongo.ListDslInfo(c, 0, 0, path, name, method, content, 0, 0, &dslList)
	if err != nil {
		logrus.Errorf("parameter invalid %v", err.Error())
		util.ResponseError(c, 500, constant.SEARCH_FAILED, "search failed")
		return
	}

	util.ResponseSuccess(c, map[string]interface{}{
		"count": count,
		"res":   dslList,
	})
}

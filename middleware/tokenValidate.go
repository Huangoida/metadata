package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"metadata/constant"
	"metadata/dal/mysql"
	"metadata/model"
	"metadata/util"
	"strconv"
)

func TokenValidate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("ApiToken")

		parseToken, err := util.ParseToken(token)
		if err != nil {
			logrus.Error("parse token failed: %v", err)
			util.ResponseError(c, 401, constant.USER_INVALID, "user unauthorized")
			c.Abort()
			return
		}
		var userList []model.UserStruct
		err, i := mysql.ListUser(c, 0, 0, "", strconv.FormatInt(parseToken.UserID, 10), &userList)
		if err != nil {
			logrus.Error("search user failed: %v", err)
			util.ResponseError(c, 500, constant.SEARCH_FAILED, "Search failed")
			c.Abort()
			return
		}
		if i != 1 {
			logrus.Error("user not found: %v", err)
			util.ResponseError(c, 401, constant.SEARCH_FAILED, "user not found")
			c.Abort()
			return
		}
		userId := strconv.FormatInt(parseToken.UserID, 10)
		c.Request.Header.Set("UserId", userId)
		c.Next()
	}
}

package util

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func ValidateOffsetAndPage(c *gin.Context) (int, int) {
	var size, page int
	sizestr := c.Query("Size")
	pageStr := c.Query("Page")
	size, err := strconv.Atoi(sizestr)
	if err != nil {
		size = 20
	}

	page, err = strconv.Atoi(pageStr)
	if err != nil {
		page = 0
	}
	return page, size
}

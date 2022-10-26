package main

import (
	"github.com/gin-gonic/gin"
	"metadata/Init"
)

func main() {
	r := gin.Default()
	Init.InitConfig()
	Init.GinRouter(r)
	r.Run()
}

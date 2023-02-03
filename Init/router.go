package Init

import (
	"metadata/handler/ApiManage"
	common "metadata/handler/Common"
	"metadata/handler/DslManage"
	"metadata/handler/Parameters"
	"metadata/handler/ServicesManage"
	"metadata/handler/UserManage"
	"metadata/middleware"

	"metadata/util"

	"github.com/gin-gonic/gin"
)

func GinRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		util.ResponseSuccess(c, "pong")
	})

	v1 := r.Group("/v1")
	services := v1.Group("/services")
	inner := v1.Group("/inner")

	services.Use(middleware.TokenValidate())
	services.POST("/create", ServicesManage.Create)
	services.GET("/list", ServicesManage.List)
	services.PUT("/update", ServicesManage.Update)
	services.DELETE("/delete", ServicesManage.Delete)

	api := v1.Group("/API")
	api.Use(middleware.TokenValidate())
	api.POST("/create", ApiManage.Create)
	api.GET("/list", ApiManage.List)
	api.PUT("/update", ApiManage.Update)
	api.DELETE("/delete", ApiManage.Delete)

	parameter := v1.Group("/parameters")
	parameter.Use(middleware.TokenValidate())
	parameter.POST("/create", Parameters.Create)
	parameter.GET("/list", Parameters.List)
	parameter.PUT("/update", Parameters.Update)
	parameter.DELETE("/delete", Parameters.Delete)

	dsl := v1.Group("/dsl")
	dsl.Use(middleware.TokenValidate())
	dsl.POST("/create", DslManage.Create)
	dsl.GET("/list", DslManage.List)
	inner.GET("/dsl/list", DslManage.InnerList)

	user := v1.Group("/user")
	user.Use(middleware.TokenValidate())
	user.POST("/create", UserManage.Create)
	user.GET("/list", UserManage.List)
	user.PUT("/update", UserManage.Update)
	user.DELETE("/delete", UserManage.Delete)

	v1.POST("/login", common.Login)
	v1.POST("/register", common.Register)

}

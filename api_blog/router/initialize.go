package router

import (
	gin "github.com/zhangdapeng520/zdpgo_gin"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_blog/router/user"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	user.InitRouter(router)

	return router
}

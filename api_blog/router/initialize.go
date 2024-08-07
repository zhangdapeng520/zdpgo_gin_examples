package router

import (
	gin "github.com/zhangdapeng520/zdpgo_gin"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_blog/router/article"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_blog/router/user"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_blog/router/wallet"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	user.InitRouter(router)
	wallet.InitRouter(router)
	article.InitRouter(router)

	return router
}

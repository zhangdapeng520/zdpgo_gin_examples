package wallet

import (
	gin "github.com/zhangdapeng520/zdpgo_gin"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_user/g"
	ginWallet "github.com/zhangdapeng520/zdpgo_gin_wallet"
)

func InitRouter(router *gin.Engine) {
	group := router.Group("/wallet")

	group.POST("/account/", ginWallet.GetAccountAddHandler(g.GDB))
	group.PUT("/account/", ginWallet.GetAccountUpdateHandler(g.GDB))
}

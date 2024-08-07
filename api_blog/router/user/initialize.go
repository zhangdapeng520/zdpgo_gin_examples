package user

import (
	gin "github.com/zhangdapeng520/zdpgo_gin"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_blog/g"
	ginLogin "github.com/zhangdapeng520/zdpgo_gin_login"
)

func InitRouter(router *gin.Engine) {
	group := router.Group("/user")

	group.POST("/register/", ginLogin.GetRegisterHandler(g.GDB, g.PasswordSalt))
	group.POST("/login/", ginLogin.GetLoginHandler(g.GDB, g.JwtKey, g.PasswordSalt))
}

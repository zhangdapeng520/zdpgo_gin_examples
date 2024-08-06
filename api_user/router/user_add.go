package router

import (
	gin "github.com/zhangdapeng520/zdpgo_gin"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_user/g"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_user/model"
)

func userAdd(c *gin.Context) {
	var user model.User
	c.ShouldBindJSON(&user)
	g.GDB.Create(&user)
	c.JSON(200, user)
}

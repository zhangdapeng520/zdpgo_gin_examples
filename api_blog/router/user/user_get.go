package user

import (
	gin "github.com/zhangdapeng520/zdpgo_gin"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_user/g"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_user/model"
)

func userGetAll(c *gin.Context) {
	var users []model.User
	g.GDB.Find(&users)
	c.JSON(200, users)
}

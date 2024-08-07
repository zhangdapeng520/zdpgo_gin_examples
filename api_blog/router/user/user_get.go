package user

import (
	gin "github.com/zhangdapeng520/zdpgo_gin"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_blog/g"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_blog/model"
)

func userGetAll(c *gin.Context) {
	var users []model.User
	g.GDB.
		Select([]string{"id", "username", "money"}).
		Find(&users)
	c.JSON(200, &users)
}

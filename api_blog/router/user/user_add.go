package user

import (
	gin "github.com/zhangdapeng520/zdpgo_gin"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_blog/g"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_blog/model"
)

func userAdd(c *gin.Context) {
	var request registerRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		g.GLog.Error(err.Error())
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	g.GDB.Create(&model.User{
		Username: request.Username,
		Password: request.Password,
	})
	c.JSON(200, nil)
}

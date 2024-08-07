package user

import (
	gin "github.com/zhangdapeng520/zdpgo_gin"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_blog/g"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_blog/model"
	password "github.com/zhangdapeng520/zdpgo_password"
)

func userAdd(c *gin.Context) {
	var request registerRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		g.GLog.Error(err.Error())
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	encryptedPassword, err := password.Sha256EncryptString(request.Password, g.Salt)
	if err != nil {
		g.GLog.Error(err.Error())
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	g.GDB.Create(&model.User{
		Username: request.Username,
		Password: encryptedPassword,
	})
	c.JSON(200, nil)
}

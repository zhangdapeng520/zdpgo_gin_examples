package user

import (
	gin "github.com/zhangdapeng520/zdpgo_gin"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_blog/g"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_blog/model"
	password "github.com/zhangdapeng520/zdpgo_password"
)

func userGetAll(c *gin.Context) {
	var users []model.User
	g.GDB.
		Select([]string{"id", "username", "money"}).
		Find(&users)
	c.JSON(200, &users)
}

func userLogin(c *gin.Context) {
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

	var user model.User
	g.GDB.Find(&model.User{
		Username: request.Username,
		Password: encryptedPassword,
	}).First(&user)

	// 生成Token
	token, err := g.GetToken(user.Id, user.Username)
	if err != nil {
		g.GLog.Error(err.Error())
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}

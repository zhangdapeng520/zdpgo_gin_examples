package router

import (
	"fmt"
	gin "github.com/zhangdapeng520/zdpgo_gin"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_user/g"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_user/model"
	"net/http"
	"strconv"
)

func userUpdate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, nil)
	}

	var reqUser model.User
	c.ShouldBind(&reqUser)

	var user = model.User{Id: id}
	g.GDB.Find(&user).First(&user)

	user.Name = reqUser.Name
	user.Age = reqUser.Age

	g.GDB.Save(&user)

	c.JSON(200, user)
}

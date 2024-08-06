package router

import (
	"fmt"
	gin "github.com/zhangdapeng520/zdpgo_gin"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_user/g"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_user/model"
	"net/http"
	"strconv"
)

func userDelete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, nil)
	}

	g.GDB.Delete(&model.User{Id: id})

	c.JSON(200, nil)
}

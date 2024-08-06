package user

import (
	gin "github.com/zhangdapeng520/zdpgo_gin"
)

func InitRouter(router *gin.Engine) {
	group := router.Group("/user")

	group.POST("/", userAdd)
	group.GET("/", userGetAll)
	group.PUT("/:id/", userUpdate)
	group.DELETE("/:id/", userDelete)
}

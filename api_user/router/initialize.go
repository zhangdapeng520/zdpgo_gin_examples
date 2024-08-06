package router

import (
	gin "github.com/zhangdapeng520/zdpgo_gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/user", userAdd)
	router.GET("/user", userGetAll)
	router.PUT("/user/:id", userUpdate)
	router.DELETE("/user/:id", userDelete)

	return router
}

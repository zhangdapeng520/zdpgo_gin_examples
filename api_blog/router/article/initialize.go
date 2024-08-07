package article

import (
	gin "github.com/zhangdapeng520/zdpgo_gin"
)

func InitRouter(router *gin.Engine) {
	group := router.Group("/article")

	group.POST("/course/", addCourse)
	group.GET("/course/", getAllCourse)
	group.PUT("/course/:id/", updateCourse)
	group.GET("/course/:id/", getCourse)
}

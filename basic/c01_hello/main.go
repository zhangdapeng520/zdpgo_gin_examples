package main

import "github.com/zhangdapeng520/zdpgo_gin"

func main() {
	router := zdpgo_gin.Default()

	router.GET("/ping", func(c *zdpgo_gin.Context) {
		c.String(200, "hello gin")
	})

	router.Run(":8080")
}

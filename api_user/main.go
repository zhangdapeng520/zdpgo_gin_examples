package main

import (
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_user/g"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_user/router"
	"net/http"
	"time"
)

func main() {
	g.InitGlobal()
	defer g.CloseGlobal()

	server := &http.Server{
		Addr:         ":8888",
		Handler:      router.InitRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	server.ListenAndServe()
}

package g

import (
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_user/model"
	gorm "github.com/zhangdapeng520/zdpgo_gorm"
	_ "github.com/zhangdapeng520/zdpgo_mysql"
)

var GDB *gorm.DB

func initMySQL() {
	var err error
	GDB, err = gorm.Open(
		"mysql",
		"root:root@tcp(127.0.0.1:3306)/test?charset=utf8",
	)
	if err != nil {
		panic(err)
	}

	GDB.DB().SetMaxIdleConns(10)
	GDB.DB().SetMaxOpenConns(100)

	GDB.AutoMigrate(&model.User{})
}

func closeMySQL() {
	GDB.Close()
}

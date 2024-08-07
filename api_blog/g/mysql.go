package g

import (
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_blog/model"
	ginLogin "github.com/zhangdapeng520/zdpgo_gin_login"
	ginWallet "github.com/zhangdapeng520/zdpgo_gin_wallet"
	gorm "github.com/zhangdapeng520/zdpgo_gorm"
	_ "github.com/zhangdapeng520/zdpgo_mysql"
)

var GDB *gorm.DB

func initMySQL() {
	var err error
	GDB, err = gorm.Open(
		"mysql",
		"root:root@tcp(127.0.0.1:3306)/blog?charset=utf8",
	)
	if err != nil {
		panic(err)
	}

	GDB.DB().SetMaxIdleConns(10)
	GDB.DB().SetMaxOpenConns(100)

	GDB.AutoMigrate(&ginLogin.GinLoginUser{})
	GDB.AutoMigrate(&ginWallet.GinWalletAccount{})
	GDB.AutoMigrate(&ginWallet.GinWalletAccountRecord{})
	GDB.AutoMigrate(&model.PythonArticle{})
	GDB.AutoMigrate(&model.GolangArticle{})
	GDB.AutoMigrate(&model.WebArticle{})
	GDB.AutoMigrate(&model.DatabaseArticle{})
	GDB.AutoMigrate(&model.CloudNativeArticle{})
	GDB.AutoMigrate(&model.CourseArticle{})
	GDB.AutoMigrate(&model.OtherArticle{})
}

func closeMySQL() {
	GDB.Close()
}

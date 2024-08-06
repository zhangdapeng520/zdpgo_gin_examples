package g

import (
	"fmt"
	_ "github.com/zhangdapeng520/zdpgo_mysql"
	zap "github.com/zhangdapeng520/zdpgo_zap"
	"time"
)

var GLog *zap.Logger

func initZap() {
	var err error
	GLog, err = zap.NewLumberjackLogger(&zap.LumberjackLoggerConfig{
		Level:      "error",
		FileName:   fmt.Sprintf("./data/log/%v.log", time.Now().Unix()),
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		IsConsole:  true,
	})
	if err != nil {
		panic(err)
	}
}

func closeZap() {
	// 调用内核的Sync方法，刷新所有缓冲的日志条目。
	// 应用程序应该注意在退出之前调用Sync。
	GLog.Sync()
}

package g

func InitGlobal() {
	initMySQL()
	initZap()
}

func CloseGlobal() {
	closeMySQL()
	closeZap()
}

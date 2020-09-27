package main

import (
	"HelloBeego190604/db_mysql"
	_ "HelloBeego190604/routers"
	"github.com/astaxie/beego"
	_ "大一下学期/github.com/go-sql-driver/mysql"
)

func main() {
	//1、连接数据库
	db_mysql.Connect()
	//2、其他配置

	//3、启动程序
beego.Run()//代码简洁
}


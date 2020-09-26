package main

import (
	_ "HelloBeego190604/routers"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_"大一下学期/github.com/go-sql-driver/mysql"
)

func main() {
	config := beego.AppConfig//定义config变量，接受并赋值为全局配置变量
	//获取配置选项
	appName := config.String("appname")
	fmt.Println("项目应用名称：",appName)
	port,err := config.Int("httpport")
	if err != nil{
		//配置信息解析错误
		panic("项目信息解析错误，请检验后重试")
	}
	fmt.Println("应用监听端口：",port)

	driver := config.String("db_driver")
	dbUser := config.String("db_root")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")
	db,err := sql.Open(driver,dbUser+":"+dbPassword+"@tcp("+dbIp+")/"+dbName+"?charset=utf8")
	//sql.Open("mysql","root:281511@tcp(127.0.0.1:3306)/hero_lol?charset=utf8")
	if err != nil{
		panic("数据库连接失败")
	}
	fmt.Println(db)
beego.Run()
}


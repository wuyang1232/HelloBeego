package controllers

import (
	"HelloBeego190604/db_mysql"
	"HelloBeego190604/models"
	"beego"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type QueryUser struct {
	beego.Controller
}
var Db *sql.DB

func Connect(){
	//项目配置
	config := beego.AppConfig //定义config变量，接受并赋值为全局配置变量
	//获取配置选项
	appName := config.String("appname")
	fmt.Println("项目应用名称：", appName)
	port, err := config.Int("httpport")
	if err != nil {
		//配置信息解析错误
		panic("项目信息解析错误，请检验后重试")
	}
	fmt.Println("应用监听端口：", port)

	driver := config.String("db_driver")
	dbUser := config.String("db_root")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")
	//1、连接数据库
	db, err := sql.Open(driver, dbUser+":"+dbPassword+"@tcp("+dbIp+")/"+dbName+"?charset=utf8")
	//sql.Open("mysql","root:281511@tcp(127.0.0.1:3306)/hero_lol?charset=utf8")
	if err != nil { //err 不等于nil表示连接数据库的时候出现错误，程序就在此中断，不用在往下执行
		//早发现，早解决
		panic("数据库连接失败") //panic：是指程序进入一种恐慌状态，程序会终止执行
	}
	Db = db
	fmt.Println(db)
	fmt.Println("数据库连接成功")
}
func (r *QueryUser) Post(){
	DataBytes,err := ioutil.ReadAll(r.Ctx.Request.Body)
	if err != nil{
		r.Ctx.WriteString("数据接收错误，请重试")
		return
	}
	var user models.Quser
	err = json.Unmarshal(DataBytes,&user)
	if err != nil{
		//r.Ctx.WriteString("数据解析错误，请重试")
		result := models.Result{
			Code:    0,
			Message: "数据解析错误，请重试",
			Data:    nil,
		}
		r.Data["json"] = &result
		r.ServeJSON()
		return
	}
	name := user.Name
	admin_num,err := db_mysql.QueryUse(name)
	if err != nil{
		//fmt.Println("123")
		fmt.Println(err.Error())
		return
	}
	if admin_num > 0{
		//md5Hash := md5.New()
		//md5Hash.Write([]byte(user.Password))
		//user.Password = hex.EncodeToString(md5Hash.Sum(nil))
		result := models.Result{
			Code:1,
			Message:"恭喜，用户注册成功",
			Data:user,
		}
		//json.Marshal(result)编码
		r.Data["json"] = &result//将result编码为json格式返回前端
		r.ServeJSON()
	}else {
		result := models.Result{
			Code:    0,
			Message: "用户名查找失败，请重试",
			Data:    nil,
		}
		r.Data["json"] = &result
		r.ServeJSON()
		return
	}
}

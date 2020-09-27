package controllers

import (
	"HelloBeego190604/db_mysql"
	"HelloBeego190604/models"
	"encoding/json"
	"fmt"
	//"beego"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Post(){
	DataBytes,err := ioutil.ReadAll(r.Ctx.Request.Body)
	if err != nil{
		r.Ctx.WriteString("数据接收错误，请重试")
		return
	}
	var user models.User
	err = json.Unmarshal(DataBytes,&user)
	if err != nil{
		r.Ctx.WriteString("数据解析错误，请重试")
		return
	}
	//一切正常，将用户信息保存到数据库中
	//直接调用保存数据的一个函数，并判断保存后的结果
	row,err := db_mysql.AddUser(user)
	if err != nil{
		r.Ctx.WriteString("注册用户信息失败，请重试")
	}
	fmt.Println("影响到的行数",row)
}

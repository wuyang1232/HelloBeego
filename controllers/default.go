package controllers

import (
	"HelloBeego190604/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller//匿名字段
}

func (c *MainController) Get() {
	//name1 := c.GetString("name")
	//age1,err := c.GetInt("age")
	//获取get请求的请求参数
	name := c.Ctx.Input.Query("name")
	age := c.Ctx.Input.Query("age")
	sex := c.Ctx.Input.Query("sex")
	fmt.Println(name,age,sex)
	//admin，18为正确数据进行验证
	if name != "admin" || age != "18"{
		c.Ctx.ResponseWriter.Write([]byte("数据验证错误"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("数据提交成功"))
	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "1403918572@qq.com"
	c.TplName = "index.tpl"
}
//该post方法是处理post类型的请求时要调用的方法
func (c *MainController) Post() {
	fmt.Println("post类型的请求...")
	user := c.Ctx.Request.FormValue("user")
	fmt.Println("用户名为：",user)
	psd := c.Ctx.Request.FormValue("psd")
	fmt.Println("密码是：",psd)

	//与固定值进行比较用户名为admin密码为123456
	if user != "admin" || psd != "123456"{
		//失败页面
		c.Ctx.ResponseWriter.Write([]byte("对不起，数据不正确"))
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("恭喜你，数据正确"))
	//request请求 response响应
	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "1403918572@qq.com"
	c.TplName = "index.tpl"

	//body := c.Ctx.Request.Body
	dataBytes,err := ioutil.ReadAll(c.Ctx.Request.Body)//获取数据
	if err != nil{
		c.Ctx.WriteString("数据接收失败，请重试")
		return
	}

	//json包解析
	var person models.Human//定义一个结构体类型数据
	err = json.Unmarshal(dataBytes,&person)//将得到的数据进行json解析到新建的结构体类型的数据里
	if err != nil{
		c.Ctx.WriteString("数据解析失败，请重试")
		return
	}
	fmt.Println("用户名：",person.Name,",年龄:",person.Age,"性别：",person.Sex)
	c.Ctx.WriteString("用户名是："+person.Name)
}
//func (c *MainController) Post(){

//}

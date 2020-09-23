package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller//匿名字段
}

func (c *MainController) Get() {
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
}

package routers

import (
	"HelloBeego190604/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/index", &controllers.MainController{})
    beego.Router("/register",&controllers.Register{})
}

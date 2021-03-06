package routers

import (
	"github.com/astaxie/beego"
	"hello/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.MainController{})
	beego.Router("/dynamics/test", &controllers.DynamicsController{})
	beego.Router("/dynamics/list", &controllers.DynamicsController{}, "get:GetAnnouceList")
	beego.Router("/dynamics/gtoken", &controllers.DynamicsController{}, "get:GenToken")
	beego.Router("/dynamics/parseform", &controllers.DynamicsController{}, "post:ParseFormData")
	beego.Router("/dynamics/postbody", &controllers.DynamicsController{}, "post:PostBody")
	beego.Router("/dynamics/postup", &controllers.DynamicsController{}, "post:PostUpload")
	beego.Router("/hmacsha/FilterOauth", &controllers.HmacshaController{}, "post:FilterOauth")
	beego.Router("/hmacsha/GetMAC", &controllers.HmacshaController{}, "post:GetMAC")

	beego.Router("/dynamics/updateFood", &controllers.DynamicsController{}, "put:UpdateFood")
	beego.Router("/dynamics/deleteFood", &controllers.DynamicsController{}, "delete:DeleteFood")
}

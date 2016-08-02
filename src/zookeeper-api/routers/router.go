package routers

import (
	"zookeeper-api/controllers"

	"github.com/astaxie/beego"
)

// Init - Init routers
func Init() {
	beego.Router("/za-faq", &controllers.FaqController{})
	ns := beego.NewNamespace("/za/v1",
		beego.NSRouter("/childs/*", &controllers.ChildsController{}),
		beego.NSRouter("/node/*", &controllers.NodeController{}),
		beego.NSRouter("/stat/:ip([0-9.:,]*)", &controllers.ZKStatController{}),
	)
	beego.AddNamespace(ns)
}

// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package common

import (
	"github.com/astaxie/beego"
	"github.com/xgopher/xbee/modules/common/controllers"
)

func init() {
	// Object RESTFUL
	beego.Router("/common/objects", &controllers.ObjectController{}, "get:Index")
	beego.Router("/common/objects", &controllers.ObjectController{}, "post:Store")
	beego.Router("/common/objects/:id", &controllers.ObjectController{}, "get:Show")
	beego.Router("/common/objects/:id", &controllers.ObjectController{}, "put:Update")
	beego.Router("/common/objects/:id", &controllers.ObjectController{}, "delete:Destroy")
}

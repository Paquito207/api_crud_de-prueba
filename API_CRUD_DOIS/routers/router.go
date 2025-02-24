// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/paquito207/API_CRUD_DOIS/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/clientes",
			beego.NSInclude(
				&controllers.ClientesController{},
			),
		),

		beego.NSNamespace("/pedidos",
			beego.NSInclude(
				&controllers.PedidosController{},
			),
		),

		beego.NSNamespace("/productos",
			beego.NSInclude(
				&controllers.ProductosController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

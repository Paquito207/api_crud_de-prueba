package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:ClientesController"] = append(beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:ClientesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:ClientesController"] = append(beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:ClientesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:ClientesController"] = append(beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:ClientesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:ClientesController"] = append(beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:ClientesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:ClientesController"] = append(beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:ClientesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:PedidosController"] = append(beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:PedidosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:PedidosController"] = append(beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:PedidosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:PedidosController"] = append(beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:PedidosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:PedidosController"] = append(beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:PedidosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:PedidosController"] = append(beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:PedidosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/paquito207/API_CRUD_UM/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

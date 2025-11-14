package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ActaRecibidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ActaRecibidoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ActaRecibidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ActaRecibidoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ActaRecibidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ActaRecibidoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ActaRecibidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ActaRecibidoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ActaRecibidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ActaRecibidoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:CampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:CampoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:CampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:CampoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:CampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:CampoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:CampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:CampoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:CampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:CampoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ElementoCampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ElementoCampoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ElementoCampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ElementoCampoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ElementoCampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ElementoCampoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ElementoCampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ElementoCampoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ElementoCampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ElementoCampoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ElementoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ElementoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ElementoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ElementoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:ElementoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:EstadoActaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:EstadoActaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:EstadoActaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:EstadoActaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:EstadoActaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:EstadoActaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:EstadoActaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:EstadoActaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:EstadoActaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:EstadoActaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:EstadoElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:EstadoElementoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:EstadoElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:EstadoElementoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:EstadoElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:EstadoElementoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:EstadoElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:EstadoElementoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:EstadoElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:EstadoElementoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:HistoricoActaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:HistoricoActaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:HistoricoActaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:HistoricoActaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:HistoricoActaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:HistoricoActaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:HistoricoActaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:HistoricoActaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:HistoricoActaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:HistoricoActaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:SoporteActaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:SoporteActaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:SoporteActaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:SoporteActaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:SoporteActaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:SoporteActaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:SoporteActaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:SoporteActaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:SoporteActaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:SoporteActaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:TipoActaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:TipoActaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:TipoActaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:TipoActaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:TipoActaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:TipoActaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:TipoActaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:TipoActaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:TipoActaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:TipoActaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:TransaccionActaRecibidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:TransaccionActaRecibidoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:TransaccionActaRecibidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:TransaccionActaRecibidoController"],
		beego.ControllerComments{
			Method:           "GetAllById",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:TransaccionActaRecibidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/acta_recibido_crud/controllers:TransaccionActaRecibidoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}

package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ActividadController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ActividadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ActividadController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ActividadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ActividadController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ActividadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ActividadController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ActividadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ActividadController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ActividadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:AreaController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:AreaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:AreaController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:AreaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:AreaController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:AreaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:AreaController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:AreaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:AreaController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:AreaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:AsignacionClienteController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:AsignacionClienteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:AsignacionClienteController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:AsignacionClienteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:AsignacionClienteController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:AsignacionClienteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:AsignacionClienteController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:AsignacionClienteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:AsignacionClienteController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:AsignacionClienteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ClienteController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ClienteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ClienteController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ClienteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ClienteController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ClienteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ClienteController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ClienteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ClienteController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ClienteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:CodigoUsuarioController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:CodigoUsuarioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:CodigoUsuarioController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:CodigoUsuarioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:CodigoUsuarioController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:CodigoUsuarioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:CodigoUsuarioController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:CodigoUsuarioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:CodigoUsuarioController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:CodigoUsuarioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:CostoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:CostoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:CostoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:CostoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:CostoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:CostoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:CostoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:CostoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:CostoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:CostoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DatosPrestacionesController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DatosPrestacionesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DatosPrestacionesController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DatosPrestacionesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DatosPrestacionesController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DatosPrestacionesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DatosPrestacionesController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DatosPrestacionesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DatosPrestacionesController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DatosPrestacionesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DescuentoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DescuentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DescuentoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DescuentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DescuentoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DescuentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DescuentoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DescuentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DescuentoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:DescuentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:PermisoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:PermisoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:PermisoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:PermisoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:PermisoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:PermisoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:PermisoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:PermisoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:PermisoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:PermisoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ProyectoPresupuestoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ProyectoPresupuestoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ProyectoPresupuestoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ProyectoPresupuestoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ProyectoPresupuestoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ProyectoPresupuestoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ProyectoPresupuestoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ProyectoPresupuestoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ProyectoPresupuestoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:ProyectoPresupuestoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:PuestoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:PuestoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:PuestoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:PuestoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:PuestoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:PuestoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:PuestoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:PuestoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:PuestoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:PuestoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:RegistroHorasController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:RegistroHorasController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:RegistroHorasController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:RegistroHorasController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:RegistroHorasController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:RegistroHorasController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:RegistroHorasController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:RegistroHorasController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:RegistroHorasController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:RegistroHorasController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:RegistroHorasPresupuestoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:RegistroHorasPresupuestoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:RegistroHorasPresupuestoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:RegistroHorasPresupuestoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:RegistroHorasPresupuestoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:RegistroHorasPresupuestoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:RegistroHorasPresupuestoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:RegistroHorasPresupuestoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:RegistroHorasPresupuestoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:RegistroHorasPresupuestoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TicketController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TicketController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TicketController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TicketController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TicketController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TicketController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TicketController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TicketController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TicketController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TicketController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TicketMessageController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TicketMessageController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TicketMessageController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TicketMessageController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TicketMessageController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TicketMessageController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TicketMessageController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TicketMessageController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TicketMessageController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TicketMessageController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TipoDescuentoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TipoDescuentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TipoDescuentoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TipoDescuentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TipoDescuentoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TipoDescuentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TipoDescuentoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TipoDescuentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TipoDescuentoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TipoDescuentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TipoPuestoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TipoPuestoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TipoPuestoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TipoPuestoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TipoPuestoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TipoPuestoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TipoPuestoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TipoPuestoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TipoPuestoController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:TipoPuestoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioSocioController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioSocioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioSocioController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioSocioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioSocioController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioSocioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioSocioController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioSocioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioSocioController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioSocioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioTrabajadorController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioTrabajadorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioTrabajadorController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioTrabajadorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioTrabajadorController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioTrabajadorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioTrabajadorController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioTrabajadorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioTrabajadorController"] = append(beego.GlobalControllerRouter["GoServerlessTest/beetest/controllers:UsuarioTrabajadorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}

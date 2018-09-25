// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"GoServerlessTest/beetest/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/actividad",
			beego.NSInclude(
				&controllers.ActividadController{},
			),
		),

		beego.NSNamespace("/area",
			beego.NSInclude(
				&controllers.AreaController{},
			),
		),

		beego.NSNamespace("/asignacion_cliente",
			beego.NSInclude(
				&controllers.AsignacionClienteController{},
			),
		),

		beego.NSNamespace("/cliente",
			beego.NSInclude(
				&controllers.ClienteController{},
			),
		),

		beego.NSNamespace("/codigo_usuario",
			beego.NSInclude(
				&controllers.CodigoUsuarioController{},
			),
		),

		beego.NSNamespace("/costo",
			beego.NSInclude(
				&controllers.CostoController{},
			),
		),

		beego.NSNamespace("/datos_prestaciones",
			beego.NSInclude(
				&controllers.DatosPrestacionesController{},
			),
		),

		beego.NSNamespace("/departamento",
			beego.NSInclude(
				&controllers.DepartamentoController{},
			),
		),

		beego.NSNamespace("/descuento",
			beego.NSInclude(
				&controllers.DescuentoController{},
			),
		),

		beego.NSNamespace("/permiso",
			beego.NSInclude(
				&controllers.PermisoController{},
			),
		),

		beego.NSNamespace("/proyecto_presupuesto",
			beego.NSInclude(
				&controllers.ProyectoPresupuestoController{},
			),
		),

		beego.NSNamespace("/puesto",
			beego.NSInclude(
				&controllers.PuestoController{},
			),
		),

		beego.NSNamespace("/registro_horas",
			beego.NSInclude(
				&controllers.RegistroHorasController{},
			),
		),

		beego.NSNamespace("/registro_horas_presupuesto",
			beego.NSInclude(
				&controllers.RegistroHorasPresupuestoController{},
			),
		),

		beego.NSNamespace("/ticket",
			beego.NSInclude(
				&controllers.TicketController{},
			),
		),

		beego.NSNamespace("/ticket_message",
			beego.NSInclude(
				&controllers.TicketMessageController{},
			),
		),

		beego.NSNamespace("/tipo_descuento",
			beego.NSInclude(
				&controllers.TipoDescuentoController{},
			),
		),

		beego.NSNamespace("/tipo_puesto",
			beego.NSInclude(
				&controllers.TipoPuestoController{},
			),
		),

		beego.NSNamespace("/usuario",
			beego.NSInclude(
				&controllers.UsuarioController{},
			),
		),

		beego.NSNamespace("/usuario_socio",
			beego.NSInclude(
				&controllers.UsuarioSocioController{},
			),
		),

		beego.NSNamespace("/usuario_trabajador",
			beego.NSInclude(
				&controllers.UsuarioTrabajadorController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

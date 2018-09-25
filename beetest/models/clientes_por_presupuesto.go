package models

type ClientesPorPresupuesto struct {
	ClienteId     *ProyectoPresupuesto `orm:"column(cliente_id);rel(fk)"`
	PresupuestoId *Cliente             `orm:"column(presupuesto_id);rel(fk)"`
}

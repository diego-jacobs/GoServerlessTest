package models

type PermisosPorTipoPuesto struct {
	TipoPuestoId *TipoPuesto `orm:"column(tipo_puesto_id);rel(fk)"`
	PermisoId    *Permiso    `orm:"column(permiso_id);rel(fk)"`
}

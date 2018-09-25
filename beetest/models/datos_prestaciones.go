package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type DatosPrestaciones struct {
	Id                      int                `orm:"column(id);auto"`
	UsuarioId               *UsuarioTrabajador `orm:"column(usuario_id);rel(fk)"`
	SueldoBase              float64            `orm:"column(sueldo_base)"`
	BonificacionIncentivo   float64            `orm:"column(bonificacion_incentivo)"`
	OtraBonificacion        float64            `orm:"column(otra_bonificacion);null"`
	Gasolina                float64            `orm:"column(gasolina);null"`
	PrestacionesSobreSueldo float64            `orm:"column(prestaciones_sobre_sueldo);null"`
	OtrasPrestaciones       float64            `orm:"column(otras_prestaciones);null"`
	Viaticos                float64            `orm:"column(viaticos);null"`
	Otros                   float64            `orm:"column(otros);null"`
	Depreciacion            float64            `orm:"column(depreciacion);null"`
	Indemnizacion           float64            `orm:"column(indemnizacion);null"`
	Aguinaldo               float64            `orm:"column(aguinaldo);null"`
	Bono14                  float64            `orm:"column(bono14);null"`
	CuotaPatronal           float64            `orm:"column(cuota_patronal);null"`
	FechaCreacion           time.Time          `orm:"column(fecha_creacion);type(datetime)"`
	FechaActualizacion      time.Time          `orm:"column(fecha_actualizacion);type(datetime)"`
	CreadoPorId             *CodigoUsuario     `orm:"column(creado_por_id);rel(fk)"`
	ActualizadoPorId        *CodigoUsuario     `orm:"column(actualizado_por_id);rel(fk)"`
	GastosIndirectos        float64            `orm:"column(gastos_indirectos);null"`
}

func (t *DatosPrestaciones) TableName() string {
	return "datos_prestaciones"
}

func init() {
	orm.RegisterModel(new(DatosPrestaciones))
}

// AddDatosPrestaciones insert a new DatosPrestaciones into database and returns
// last inserted Id on success.
func AddDatosPrestaciones(m *DatosPrestaciones) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDatosPrestacionesById retrieves DatosPrestaciones by Id. Returns error if
// Id doesn't exist
func GetDatosPrestacionesById(id int) (v *DatosPrestaciones, err error) {
	o := orm.NewOrm()
	v = &DatosPrestaciones{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDatosPrestaciones retrieves all DatosPrestaciones matches certain condition. Returns empty list if
// no records exist
func GetAllDatosPrestaciones(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DatosPrestaciones))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []DatosPrestaciones
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateDatosPrestaciones updates DatosPrestaciones by Id and returns error if
// the record to be updated doesn't exist
func UpdateDatosPrestacionesById(m *DatosPrestaciones) (err error) {
	o := orm.NewOrm()
	v := DatosPrestaciones{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDatosPrestaciones deletes DatosPrestaciones by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDatosPrestaciones(id int) (err error) {
	o := orm.NewOrm()
	v := DatosPrestaciones{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DatosPrestaciones{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

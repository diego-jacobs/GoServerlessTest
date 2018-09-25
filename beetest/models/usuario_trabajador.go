package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type UsuarioTrabajador struct {
	Id                   int       `orm:"column(id);pk"`
	Direccion            string    `orm:"column(direccion);size(255);null"`
	FechaEgreso          time.Time `orm:"column(fecha_egreso);type(date);null"`
	Dpi                  string    `orm:"column(dpi);size(20);null"`
	Nit                  string    `orm:"column(nit);size(20);null"`
	Telefono             string    `orm:"column(telefono);size(15);null"`
	NumeroAfiliacionIgss string    `orm:"column(numero_afiliacion_igss);size(15);null"`
}

func (t *UsuarioTrabajador) TableName() string {
	return "usuario_trabajador"
}

func init() {
	orm.RegisterModel(new(UsuarioTrabajador))
}

// AddUsuarioTrabajador insert a new UsuarioTrabajador into database and returns
// last inserted Id on success.
func AddUsuarioTrabajador(m *UsuarioTrabajador) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUsuarioTrabajadorById retrieves UsuarioTrabajador by Id. Returns error if
// Id doesn't exist
func GetUsuarioTrabajadorById(id int) (v *UsuarioTrabajador, err error) {
	o := orm.NewOrm()
	v = &UsuarioTrabajador{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUsuarioTrabajador retrieves all UsuarioTrabajador matches certain condition. Returns empty list if
// no records exist
func GetAllUsuarioTrabajador(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UsuarioTrabajador))
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

	var l []UsuarioTrabajador
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

// UpdateUsuarioTrabajador updates UsuarioTrabajador by Id and returns error if
// the record to be updated doesn't exist
func UpdateUsuarioTrabajadorById(m *UsuarioTrabajador) (err error) {
	o := orm.NewOrm()
	v := UsuarioTrabajador{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUsuarioTrabajador deletes UsuarioTrabajador by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUsuarioTrabajador(id int) (err error) {
	o := orm.NewOrm()
	v := UsuarioTrabajador{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&UsuarioTrabajador{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// Package models contains the types for schema 'v1wq1ics1m037sn6'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
)

// PermisosPorTipoPuesto represents a row from 'v1wq1ics1m037sn6.permisos_por_tipo_puesto'.
type PermisosPorTipoPuesto struct {
	TipoPuestoID int `json:"tipo_puesto_id"` // tipo_puesto_id
	PermisoID    int `json:"permiso_id"`     // permiso_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the PermisosPorTipoPuesto exists in the database.
func (pptp *PermisosPorTipoPuesto) Exists() bool {
	return pptp._exists
}

// Deleted provides information if the PermisosPorTipoPuesto has been deleted from the database.
func (pptp *PermisosPorTipoPuesto) Deleted() bool {
	return pptp._deleted
}

// Insert inserts the PermisosPorTipoPuesto to the database.
func (pptp *PermisosPorTipoPuesto) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if pptp._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO v1wq1ics1m037sn6.permisos_por_tipo_puesto (` +
		`tipo_puesto_id, permiso_id` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
	XOLog(sqlstr, pptp.TipoPuestoID, pptp.PermisoID)
	_, err = db.Exec(sqlstr, pptp.TipoPuestoID, pptp.PermisoID)
	if err != nil {
		return err
	}

	// set existence
	pptp._exists = true

	return nil
}

// Update statements omitted due to lack of fields other than primary key

// Delete deletes the PermisosPorTipoPuesto from the database.
func (pptp *PermisosPorTipoPuesto) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !pptp._exists {
		return nil
	}

	// if deleted, bail
	if pptp._deleted {
		return nil
	}

	// sql query with composite primary key
	const sqlstr = `DELETE FROM v1wq1ics1m037sn6.permisos_por_tipo_puesto WHERE tipo_puesto_id = ? AND permiso_id = ?`

	// run query
	XOLog(sqlstr, pptp.TipoPuestoID, pptp.PermisoID)
	_, err = db.Exec(sqlstr, pptp.TipoPuestoID, pptp.PermisoID)
	if err != nil {
		return err
	}

	// set deleted
	pptp._deleted = true

	return nil
}

// Permiso returns the Permiso associated with the PermisosPorTipoPuesto's PermisoID (permiso_id).
//
// Generated from foreign key 'FK_B882F8E6CEFAD37'.
func (pptp *PermisosPorTipoPuesto) Permiso(db XODB) (*Permiso, error) {
	return PermisoByID(db, pptp.PermisoID)
}

// TipoPuesto returns the TipoPuesto associated with the PermisosPorTipoPuesto's TipoPuestoID (tipo_puesto_id).
//
// Generated from foreign key 'FK_B882F8EB25B5694'.
func (pptp *PermisosPorTipoPuesto) TipoPuesto(db XODB) (*TipoPuesto, error) {
	return TipoPuestoByID(db, pptp.TipoPuestoID)
}

// PermisosPorTipoPuestosByPermisoID retrieves a row from 'v1wq1ics1m037sn6.permisos_por_tipo_puesto' as a PermisosPorTipoPuesto.
//
// Generated from index 'IDX_B882F8E6CEFAD37'.
func PermisosPorTipoPuestosByPermisoID(db XODB, permisoID int) ([]*PermisosPorTipoPuesto, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`tipo_puesto_id, permiso_id ` +
		`FROM v1wq1ics1m037sn6.permisos_por_tipo_puesto ` +
		`WHERE permiso_id = ?`

	// run query
	XOLog(sqlstr, permisoID)
	q, err := db.Query(sqlstr, permisoID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*PermisosPorTipoPuesto{}
	for q.Next() {
		pptp := PermisosPorTipoPuesto{
			_exists: true,
		}

		// scan
		err = q.Scan(&pptp.TipoPuestoID, &pptp.PermisoID)
		if err != nil {
			return nil, err
		}

		res = append(res, &pptp)
	}

	return res, nil
}

// PermisosPorTipoPuestosByTipoPuestoID retrieves a row from 'v1wq1ics1m037sn6.permisos_por_tipo_puesto' as a PermisosPorTipoPuesto.
//
// Generated from index 'IDX_B882F8EB25B5694'.
func PermisosPorTipoPuestosByTipoPuestoID(db XODB, tipoPuestoID int) ([]*PermisosPorTipoPuesto, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`tipo_puesto_id, permiso_id ` +
		`FROM v1wq1ics1m037sn6.permisos_por_tipo_puesto ` +
		`WHERE tipo_puesto_id = ?`

	// run query
	XOLog(sqlstr, tipoPuestoID)
	q, err := db.Query(sqlstr, tipoPuestoID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*PermisosPorTipoPuesto{}
	for q.Next() {
		pptp := PermisosPorTipoPuesto{
			_exists: true,
		}

		// scan
		err = q.Scan(&pptp.TipoPuestoID, &pptp.PermisoID)
		if err != nil {
			return nil, err
		}

		res = append(res, &pptp)
	}

	return res, nil
}

// PermisosPorTipoPuestoByPermisoID retrieves a row from 'v1wq1ics1m037sn6.permisos_por_tipo_puesto' as a PermisosPorTipoPuesto.
//
// Generated from index 'permisos_por_tipo_puesto_permiso_id_pkey'.
func PermisosPorTipoPuestoByPermisoID(db XODB, permisoID int) (*PermisosPorTipoPuesto, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`tipo_puesto_id, permiso_id ` +
		`FROM v1wq1ics1m037sn6.permisos_por_tipo_puesto ` +
		`WHERE permiso_id = ?`

	// run query
	XOLog(sqlstr, permisoID)
	pptp := PermisosPorTipoPuesto{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, permisoID).Scan(&pptp.TipoPuestoID, &pptp.PermisoID)
	if err != nil {
		return nil, err
	}

	return &pptp, nil
}

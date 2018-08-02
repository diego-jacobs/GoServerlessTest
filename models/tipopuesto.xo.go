// Package models contains the types for schema 'v1wq1ics1m037sn6'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// TipoPuesto represents a row from 'v1wq1ics1m037sn6.tipo_puesto'.
type TipoPuesto struct {
	ID                 int            `json:"id"`                  // id
	NombrePuesto       string         `json:"nombre_puesto"`       // nombre_puesto
	Descripcion        sql.NullString `json:"descripcion"`         // descripcion
	Abreviatura        sql.NullString `json:"abreviatura"`         // abreviatura
	CreadoPorID        sql.NullInt64  `json:"creado_por_id"`       // creado_por_id
	ActualizadoPorID   sql.NullInt64  `json:"actualizado_por_id"`  // actualizado_por_id
	FechaCreacion      time.Time      `json:"fecha_creacion"`      // fecha_creacion
	FechaActualizacion time.Time      `json:"fecha_actualizacion"` // fecha_actualizacion

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the TipoPuesto exists in the database.
func (tp *TipoPuesto) Exists() bool {
	return tp._exists
}

// Deleted provides information if the TipoPuesto has been deleted from the database.
func (tp *TipoPuesto) Deleted() bool {
	return tp._deleted
}

// Insert inserts the TipoPuesto to the database.
func (tp *TipoPuesto) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if tp._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO v1wq1ics1m037sn6.tipo_puesto (` +
		`nombre_puesto, descripcion, abreviatura, creado_por_id, actualizado_por_id, fecha_creacion, fecha_actualizacion` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, tp.NombrePuesto, tp.Descripcion, tp.Abreviatura, tp.CreadoPorID, tp.ActualizadoPorID, tp.FechaCreacion, tp.FechaActualizacion)
	res, err := db.Exec(sqlstr, tp.NombrePuesto, tp.Descripcion, tp.Abreviatura, tp.CreadoPorID, tp.ActualizadoPorID, tp.FechaCreacion, tp.FechaActualizacion)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	tp.ID = int(id)
	tp._exists = true

	return nil
}

// Update updates the TipoPuesto in the database.
func (tp *TipoPuesto) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !tp._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if tp._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE v1wq1ics1m037sn6.tipo_puesto SET ` +
		`nombre_puesto = ?, descripcion = ?, abreviatura = ?, creado_por_id = ?, actualizado_por_id = ?, fecha_creacion = ?, fecha_actualizacion = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, tp.NombrePuesto, tp.Descripcion, tp.Abreviatura, tp.CreadoPorID, tp.ActualizadoPorID, tp.FechaCreacion, tp.FechaActualizacion, tp.ID)
	_, err = db.Exec(sqlstr, tp.NombrePuesto, tp.Descripcion, tp.Abreviatura, tp.CreadoPorID, tp.ActualizadoPorID, tp.FechaCreacion, tp.FechaActualizacion, tp.ID)
	return err
}

// Save saves the TipoPuesto to the database.
func (tp *TipoPuesto) Save(db XODB) error {
	if tp.Exists() {
		return tp.Update(db)
	}

	return tp.Insert(db)
}

// Delete deletes the TipoPuesto from the database.
func (tp *TipoPuesto) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !tp._exists {
		return nil
	}

	// if deleted, bail
	if tp._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM v1wq1ics1m037sn6.tipo_puesto WHERE id = ?`

	// run query
	XOLog(sqlstr, tp.ID)
	_, err = db.Exec(sqlstr, tp.ID)
	if err != nil {
		return err
	}

	// set deleted
	tp._deleted = true

	return nil
}

// CodigoUsuarioByActualizadoPorID returns the CodigoUsuario associated with the TipoPuesto's ActualizadoPorID (actualizado_por_id).
//
// Generated from foreign key 'FK_D27A24C6F6167A1C'.
func (tp *TipoPuesto) CodigoUsuarioByActualizadoPorID(db XODB) (*CodigoUsuario, error) {
	return CodigoUsuarioByID(db, int(tp.ActualizadoPorID.Int64))
}

// CodigoUsuarioByCreadoPorID returns the CodigoUsuario associated with the TipoPuesto's CreadoPorID (creado_por_id).
//
// Generated from foreign key 'FK_D27A24C6FE35D8C4'.
func (tp *TipoPuesto) CodigoUsuarioByCreadoPorID(db XODB) (*CodigoUsuario, error) {
	return CodigoUsuarioByID(db, int(tp.CreadoPorID.Int64))
}

// TipoPuestosByActualizadoPorID retrieves a row from 'v1wq1ics1m037sn6.tipo_puesto' as a TipoPuesto.
//
// Generated from index 'IDX_D27A24C6F6167A1C'.
func TipoPuestosByActualizadoPorID(db XODB, actualizadoPorID sql.NullInt64) ([]*TipoPuesto, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, nombre_puesto, descripcion, abreviatura, creado_por_id, actualizado_por_id, fecha_creacion, fecha_actualizacion ` +
		`FROM v1wq1ics1m037sn6.tipo_puesto ` +
		`WHERE actualizado_por_id = ?`

	// run query
	XOLog(sqlstr, actualizadoPorID)
	q, err := db.Query(sqlstr, actualizadoPorID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*TipoPuesto{}
	for q.Next() {
		tp := TipoPuesto{
			_exists: true,
		}

		// scan
		err = q.Scan(&tp.ID, &tp.NombrePuesto, &tp.Descripcion, &tp.Abreviatura, &tp.CreadoPorID, &tp.ActualizadoPorID, &tp.FechaCreacion, &tp.FechaActualizacion)
		if err != nil {
			return nil, err
		}

		res = append(res, &tp)
	}

	return res, nil
}

// TipoPuestosByCreadoPorID retrieves a row from 'v1wq1ics1m037sn6.tipo_puesto' as a TipoPuesto.
//
// Generated from index 'IDX_D27A24C6FE35D8C4'.
func TipoPuestosByCreadoPorID(db XODB, creadoPorID sql.NullInt64) ([]*TipoPuesto, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, nombre_puesto, descripcion, abreviatura, creado_por_id, actualizado_por_id, fecha_creacion, fecha_actualizacion ` +
		`FROM v1wq1ics1m037sn6.tipo_puesto ` +
		`WHERE creado_por_id = ?`

	// run query
	XOLog(sqlstr, creadoPorID)
	q, err := db.Query(sqlstr, creadoPorID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*TipoPuesto{}
	for q.Next() {
		tp := TipoPuesto{
			_exists: true,
		}

		// scan
		err = q.Scan(&tp.ID, &tp.NombrePuesto, &tp.Descripcion, &tp.Abreviatura, &tp.CreadoPorID, &tp.ActualizadoPorID, &tp.FechaCreacion, &tp.FechaActualizacion)
		if err != nil {
			return nil, err
		}

		res = append(res, &tp)
	}

	return res, nil
}

// TipoPuestoByID retrieves a row from 'v1wq1ics1m037sn6.tipo_puesto' as a TipoPuesto.
//
// Generated from index 'tipo_puesto_id_pkey'.
func TipoPuestoByID(db XODB, id int) (*TipoPuesto, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, nombre_puesto, descripcion, abreviatura, creado_por_id, actualizado_por_id, fecha_creacion, fecha_actualizacion ` +
		`FROM v1wq1ics1m037sn6.tipo_puesto ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	tp := TipoPuesto{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&tp.ID, &tp.NombrePuesto, &tp.Descripcion, &tp.Abreviatura, &tp.CreadoPorID, &tp.ActualizadoPorID, &tp.FechaCreacion, &tp.FechaActualizacion)
	if err != nil {
		return nil, err
	}

	return &tp, nil
}
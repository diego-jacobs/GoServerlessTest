// Package models contains the types for schema 'v1wq1ics1m037sn6'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// Departamento represents a row from 'v1wq1ics1m037sn6.departamento'.
type Departamento struct {
	ID                 int            `json:"id"`                  // id
	NombreDepartamento string         `json:"nombre_departamento"` // nombre_departamento
	Descripcion        sql.NullString `json:"descripcion"`         // descripcion
	CreadoPorID        sql.NullInt64  `json:"creado_por_id"`       // creado_por_id
	ActualizadoPorID   sql.NullInt64  `json:"actualizado_por_id"`  // actualizado_por_id
	FechaCreacion      time.Time      `json:"fecha_creacion"`      // fecha_creacion
	FechaActualizacion time.Time      `json:"fecha_actualizacion"` // fecha_actualizacion

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Departamento exists in the database.
func (d *Departamento) Exists() bool {
	return d._exists
}

// Deleted provides information if the Departamento has been deleted from the database.
func (d *Departamento) Deleted() bool {
	return d._deleted
}

// Insert inserts the Departamento to the database.
func (d *Departamento) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if d._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO v1wq1ics1m037sn6.departamento (` +
		`nombre_departamento, descripcion, creado_por_id, actualizado_por_id, fecha_creacion, fecha_actualizacion` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, d.NombreDepartamento, d.Descripcion, d.CreadoPorID, d.ActualizadoPorID, d.FechaCreacion, d.FechaActualizacion)
	res, err := db.Exec(sqlstr, d.NombreDepartamento, d.Descripcion, d.CreadoPorID, d.ActualizadoPorID, d.FechaCreacion, d.FechaActualizacion)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	d.ID = int(id)
	d._exists = true

	return nil
}

// Update updates the Departamento in the database.
func (d *Departamento) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !d._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if d._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE v1wq1ics1m037sn6.departamento SET ` +
		`nombre_departamento = ?, descripcion = ?, creado_por_id = ?, actualizado_por_id = ?, fecha_creacion = ?, fecha_actualizacion = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, d.NombreDepartamento, d.Descripcion, d.CreadoPorID, d.ActualizadoPorID, d.FechaCreacion, d.FechaActualizacion, d.ID)
	_, err = db.Exec(sqlstr, d.NombreDepartamento, d.Descripcion, d.CreadoPorID, d.ActualizadoPorID, d.FechaCreacion, d.FechaActualizacion, d.ID)
	return err
}

// Save saves the Departamento to the database.
func (d *Departamento) Save(db XODB) error {
	if d.Exists() {
		return d.Update(db)
	}

	return d.Insert(db)
}

// Delete deletes the Departamento from the database.
func (d *Departamento) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !d._exists {
		return nil
	}

	// if deleted, bail
	if d._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM v1wq1ics1m037sn6.departamento WHERE id = ?`

	// run query
	XOLog(sqlstr, d.ID)
	_, err = db.Exec(sqlstr, d.ID)
	if err != nil {
		return err
	}

	// set deleted
	d._deleted = true

	return nil
}

// CodigoUsuarioByActualizadoPorID returns the CodigoUsuario associated with the Departamento's ActualizadoPorID (actualizado_por_id).
//
// Generated from foreign key 'FK_40E497EBF6167A1C'.
func (d *Departamento) CodigoUsuarioByActualizadoPorID(db XODB) (*CodigoUsuario, error) {
	return CodigoUsuarioByID(db, int(d.ActualizadoPorID.Int64))
}

// CodigoUsuarioByCreadoPorID returns the CodigoUsuario associated with the Departamento's CreadoPorID (creado_por_id).
//
// Generated from foreign key 'FK_40E497EBFE35D8C4'.
func (d *Departamento) CodigoUsuarioByCreadoPorID(db XODB) (*CodigoUsuario, error) {
	return CodigoUsuarioByID(db, int(d.CreadoPorID.Int64))
}

// DepartamentosByActualizadoPorID retrieves a row from 'v1wq1ics1m037sn6.departamento' as a Departamento.
//
// Generated from index 'IDX_40E497EBF6167A1C'.
func DepartamentosByActualizadoPorID(db XODB, actualizadoPorID sql.NullInt64) ([]*Departamento, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, nombre_departamento, descripcion, creado_por_id, actualizado_por_id, fecha_creacion, fecha_actualizacion ` +
		`FROM v1wq1ics1m037sn6.departamento ` +
		`WHERE actualizado_por_id = ?`

	// run query
	XOLog(sqlstr, actualizadoPorID)
	q, err := db.Query(sqlstr, actualizadoPorID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Departamento{}
	for q.Next() {
		d := Departamento{
			_exists: true,
		}

		// scan
		err = q.Scan(&d.ID, &d.NombreDepartamento, &d.Descripcion, &d.CreadoPorID, &d.ActualizadoPorID, &d.FechaCreacion, &d.FechaActualizacion)
		if err != nil {
			return nil, err
		}

		res = append(res, &d)
	}

	return res, nil
}

// DepartamentosByCreadoPorID retrieves a row from 'v1wq1ics1m037sn6.departamento' as a Departamento.
//
// Generated from index 'IDX_40E497EBFE35D8C4'.
func DepartamentosByCreadoPorID(db XODB, creadoPorID sql.NullInt64) ([]*Departamento, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, nombre_departamento, descripcion, creado_por_id, actualizado_por_id, fecha_creacion, fecha_actualizacion ` +
		`FROM v1wq1ics1m037sn6.departamento ` +
		`WHERE creado_por_id = ?`

	// run query
	XOLog(sqlstr, creadoPorID)
	q, err := db.Query(sqlstr, creadoPorID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Departamento{}
	for q.Next() {
		d := Departamento{
			_exists: true,
		}

		// scan
		err = q.Scan(&d.ID, &d.NombreDepartamento, &d.Descripcion, &d.CreadoPorID, &d.ActualizadoPorID, &d.FechaCreacion, &d.FechaActualizacion)
		if err != nil {
			return nil, err
		}

		res = append(res, &d)
	}

	return res, nil
}

// DepartamentoByID retrieves a row from 'v1wq1ics1m037sn6.departamento' as a Departamento.
//
// Generated from index 'departamento_id_pkey'.
func DepartamentoByID(db XODB, id int) (*Departamento, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, nombre_departamento, descripcion, creado_por_id, actualizado_por_id, fecha_creacion, fecha_actualizacion ` +
		`FROM v1wq1ics1m037sn6.departamento ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	d := Departamento{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&d.ID, &d.NombreDepartamento, &d.Descripcion, &d.CreadoPorID, &d.ActualizadoPorID, &d.FechaCreacion, &d.FechaActualizacion)
	if err != nil {
		return nil, err
	}

	return &d, nil
}

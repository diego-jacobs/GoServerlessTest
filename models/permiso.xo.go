// Package models contains the types for schema 'v1wq1ics1m037sn6'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
	"time"
)

// Permiso represents a row from 'v1wq1ics1m037sn6.permiso'.
type Permiso struct {
	ID                 int       `json:"id"`                  // id
	Etiqueta           string    `json:"etiqueta"`            // etiqueta
	Permiso            string    `json:"permiso"`             // permiso
	FechaCreacion      time.Time `json:"fecha_creacion"`      // fecha_creacion
	FechaActualizacion time.Time `json:"fecha_actualizacion"` // fecha_actualizacion
	CreadoPor          string    `json:"creado_por"`          // creado_por
	ActualizadoPor     string    `json:"actualizado_por"`     // actualizado_por

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Permiso exists in the database.
func (p *Permiso) Exists() bool {
	return p._exists
}

// Deleted provides information if the Permiso has been deleted from the database.
func (p *Permiso) Deleted() bool {
	return p._deleted
}

// Insert inserts the Permiso to the database.
func (p *Permiso) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if p._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO v1wq1ics1m037sn6.permiso (` +
		`etiqueta, permiso, fecha_creacion, fecha_actualizacion, creado_por, actualizado_por` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, p.Etiqueta, p.Permiso, p.FechaCreacion, p.FechaActualizacion, p.CreadoPor, p.ActualizadoPor)
	res, err := db.Exec(sqlstr, p.Etiqueta, p.Permiso, p.FechaCreacion, p.FechaActualizacion, p.CreadoPor, p.ActualizadoPor)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	p.ID = int(id)
	p._exists = true

	return nil
}

// Update updates the Permiso in the database.
func (p *Permiso) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !p._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if p._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE v1wq1ics1m037sn6.permiso SET ` +
		`etiqueta = ?, permiso = ?, fecha_creacion = ?, fecha_actualizacion = ?, creado_por = ?, actualizado_por = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, p.Etiqueta, p.Permiso, p.FechaCreacion, p.FechaActualizacion, p.CreadoPor, p.ActualizadoPor, p.ID)
	_, err = db.Exec(sqlstr, p.Etiqueta, p.Permiso, p.FechaCreacion, p.FechaActualizacion, p.CreadoPor, p.ActualizadoPor, p.ID)
	return err
}

// Save saves the Permiso to the database.
func (p *Permiso) Save(db XODB) error {
	if p.Exists() {
		return p.Update(db)
	}

	return p.Insert(db)
}

// Delete deletes the Permiso from the database.
func (p *Permiso) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !p._exists {
		return nil
	}

	// if deleted, bail
	if p._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM v1wq1ics1m037sn6.permiso WHERE id = ?`

	// run query
	XOLog(sqlstr, p.ID)
	_, err = db.Exec(sqlstr, p.ID)
	if err != nil {
		return err
	}

	// set deleted
	p._deleted = true

	return nil
}

// PermisoByID retrieves a row from 'v1wq1ics1m037sn6.permiso' as a Permiso.
//
// Generated from index 'permiso_id_pkey'.
func PermisoByID(db XODB, id int) (*Permiso, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, etiqueta, permiso, fecha_creacion, fecha_actualizacion, creado_por, actualizado_por ` +
		`FROM v1wq1ics1m037sn6.permiso ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	p := Permiso{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&p.ID, &p.Etiqueta, &p.Permiso, &p.FechaCreacion, &p.FechaActualizacion, &p.CreadoPor, &p.ActualizadoPor)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

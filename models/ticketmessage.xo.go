// Package models contains the types for schema 'v1wq1ics1m037sn6'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// TicketMessage represents a row from 'v1wq1ics1m037sn6.ticket_message'.
type TicketMessage struct {
	ID        int            `json:"id"`         // id
	TicketID  sql.NullInt64  `json:"ticket_id"`  // ticket_id
	UserID    int            `json:"user_id"`    // user_id
	Message   sql.NullString `json:"message"`    // message
	Status    int16          `json:"status"`     // status
	Priority  int16          `json:"priority"`   // priority
	CreatedAt time.Time      `json:"created_at"` // created_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the TicketMessage exists in the database.
func (tm *TicketMessage) Exists() bool {
	return tm._exists
}

// Deleted provides information if the TicketMessage has been deleted from the database.
func (tm *TicketMessage) Deleted() bool {
	return tm._deleted
}

// Insert inserts the TicketMessage to the database.
func (tm *TicketMessage) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if tm._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO v1wq1ics1m037sn6.ticket_message (` +
		`ticket_id, user_id, message, status, priority, created_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, tm.TicketID, tm.UserID, tm.Message, tm.Status, tm.Priority, tm.CreatedAt)
	res, err := db.Exec(sqlstr, tm.TicketID, tm.UserID, tm.Message, tm.Status, tm.Priority, tm.CreatedAt)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	tm.ID = int(id)
	tm._exists = true

	return nil
}

// Update updates the TicketMessage in the database.
func (tm *TicketMessage) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !tm._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if tm._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE v1wq1ics1m037sn6.ticket_message SET ` +
		`ticket_id = ?, user_id = ?, message = ?, status = ?, priority = ?, created_at = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, tm.TicketID, tm.UserID, tm.Message, tm.Status, tm.Priority, tm.CreatedAt, tm.ID)
	_, err = db.Exec(sqlstr, tm.TicketID, tm.UserID, tm.Message, tm.Status, tm.Priority, tm.CreatedAt, tm.ID)
	return err
}

// Save saves the TicketMessage to the database.
func (tm *TicketMessage) Save(db XODB) error {
	if tm.Exists() {
		return tm.Update(db)
	}

	return tm.Insert(db)
}

// Delete deletes the TicketMessage from the database.
func (tm *TicketMessage) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !tm._exists {
		return nil
	}

	// if deleted, bail
	if tm._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM v1wq1ics1m037sn6.ticket_message WHERE id = ?`

	// run query
	XOLog(sqlstr, tm.ID)
	_, err = db.Exec(sqlstr, tm.ID)
	if err != nil {
		return err
	}

	// set deleted
	tm._deleted = true

	return nil
}

// Ticket returns the Ticket associated with the TicketMessage's TicketID (ticket_id).
//
// Generated from foreign key 'FK_BA71692D700047D2'.
func (tm *TicketMessage) Ticket(db XODB) (*Ticket, error) {
	return TicketByID(db, int(tm.TicketID.Int64))
}

// TicketMessagesByTicketID retrieves a row from 'v1wq1ics1m037sn6.ticket_message' as a TicketMessage.
//
// Generated from index 'IDX_BA71692D700047D2'.
func TicketMessagesByTicketID(db XODB, ticketID sql.NullInt64) ([]*TicketMessage, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, ticket_id, user_id, message, status, priority, created_at ` +
		`FROM v1wq1ics1m037sn6.ticket_message ` +
		`WHERE ticket_id = ?`

	// run query
	XOLog(sqlstr, ticketID)
	q, err := db.Query(sqlstr, ticketID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*TicketMessage{}
	for q.Next() {
		tm := TicketMessage{
			_exists: true,
		}

		// scan
		err = q.Scan(&tm.ID, &tm.TicketID, &tm.UserID, &tm.Message, &tm.Status, &tm.Priority, &tm.CreatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &tm)
	}

	return res, nil
}

// TicketMessageByID retrieves a row from 'v1wq1ics1m037sn6.ticket_message' as a TicketMessage.
//
// Generated from index 'ticket_message_id_pkey'.
func TicketMessageByID(db XODB, id int) (*TicketMessage, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, ticket_id, user_id, message, status, priority, created_at ` +
		`FROM v1wq1ics1m037sn6.ticket_message ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	tm := TicketMessage{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&tm.ID, &tm.TicketID, &tm.UserID, &tm.Message, &tm.Status, &tm.Priority, &tm.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &tm, nil
}
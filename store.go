package rethink

import (
	"time"

	r "github.com/dancannon/gorethink"
)

// DB represents a connection to the DB and allows you to run queriedb.
type DB struct {
	Session     *r.Session
	collections *[]Collection
}

// Connect establishes connection with rethinkDB
func Connect(address, database string) (DB, error) {

	var err error
	db := DB{}
	db.Session, err = r.Connect(r.ConnectOpts{
		Address:     address,
		Database:    database,
		MaxIdle:     10,
		IdleTimeout: time.Second * 10,
	})
	if err != nil {
		return db, err
	}

	return db, nil
}

// NewFromSession returns a new DB from an existing gorethink session.
func NewFromSession(session *r.Session) DB {
	return DB{Session: session}
}

// NewCollection returns a new collection with the db set.
func (db *DB) NewCollection(name string) *Collection {
	return &Collection{r.Table(name), db}
}

// RunWrite will run a query for the current session.
func (db *DB) RunWrite(term r.Term) (r.WriteResponse, error) {
	writeRes := r.WriteResponse{}
	writeRes, err := term.RunWrite(db.Session)
	if err != nil {
		return writeRes, err
	}

	return writeRes, nil
}

// Run will run a query and return the cursor.
func (db *DB) Run(term r.Term) (*r.Cursor, error) {
	cursor, err := term.Run(db.Session)
	if err != nil {
		return nil, err
	}

	return cursor, nil
}

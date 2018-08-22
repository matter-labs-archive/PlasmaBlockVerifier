package database

import (
	"github.com/dgraph-io/badger"
)

var database *badger.DB

// Open a Badger database in a local ./db directory
func OpenDatabase() (*badger.DB, error) {
	if database != nil {
		return database, nil
	}
	opts := badger.DefaultOptions
	opts.Dir = "./db"
	opts.ValueDir = "./db"
	db, err := badger.Open(opts)

	if err != nil {
		return nil, err
	}
	database = db
	return database, nil
}

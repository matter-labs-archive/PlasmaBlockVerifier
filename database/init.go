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

// Open a test Badger database in a /tmp/badger directory
func OpenTestDatabase() (*badger.DB, error) {
	opts := badger.DefaultOptions
	opts.Dir = "/tmp/badger"
	opts.ValueDir = "/tmp/badger"
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	return db, nil
}

package database

import (
	"github.com/asdine/storm"
)

var (
	db *storm.DB
)

func Open(name string) (err error) {
	db, err = storm.Open(name)
	return err
}

func Close() error {
	return db.Close()
}

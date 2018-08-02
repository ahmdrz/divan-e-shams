package database

import (
	"github.com/asdine/storm/q"
)

type Poem struct {
	ID      uint   `storm:"id"`
	Content string `storm:"index"`
}

func (p *Poem) Save() error {
	return db.Save(p)
}

func GetPoem(selector string, value interface{}) (*Poem, error) {
	poem := &Poem{}
	err := db.One(selector, value, poem)
	return poem, err
}

func FindPoem(query string) (poem []Poem, err error) {
	err = db.Select(q.Re("Content", query)).Find(&poem)
	return
}

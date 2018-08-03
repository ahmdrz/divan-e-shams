package database

import (
	"github.com/asdine/storm/q"
)

type Robaei struct {
	ID      uint   `storm:"id"`
	Content string `storm:"index"`
}

func (p *Robaei) Save() error {
	return db.Save(p)
}

func GetRobaei(selector string, value interface{}) (*Robaei, error) {
	robaei := &Robaei{}
	err := db.One(selector, value, robaei)
	return robaei, err
}

func FindRobaei(query string) (robaei []Robaei, err error) {
	err = db.Select(q.Re("Content", query)).Find(&robaei)
	return
}

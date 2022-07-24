package biz

import (
	"database/sql"
)

type Obj struct {
	db *sql.DB
}

func NewObj(s *sql.DB) *Obj{
	return &Obj{db: s}
}

func (o *Obj)Get(key string)(value interface{}){
	//...
	return nil
}
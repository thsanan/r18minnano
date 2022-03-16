package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/thsanan/r18minnano/pogo"
)

type ActressRepository interface {
	GetById(actId string) (pogo.Actress, error)
	SelectAll() ([]pogo.Actress, error)
}

type ActressMySql struct {
	db *sqlx.DB
}

func (act ActressMySql) GetById(actId string) (pogo.Actress, error) {
	//TODO implement me
	panic("implement me")
}

func (act ActressMySql) SelectAll() ([]pogo.Actress, error) {
	//TODO implement me
	panic("implement me")
}

func NewActressMysql(db *sqlx.DB) ActressRepository {
	return ActressMySql{db}
}

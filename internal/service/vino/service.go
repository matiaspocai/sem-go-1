package vino

import (
	"github.com/jmoiron/sqlx"
	"github.com/matiaspocai/sem-go-1/internal/config"
)

// Message ...
type Vino struct {
	ID   int
	Text string
}

// Service ...
type Service interface {
	AddVino(Vino) error
	FindAll() []*Vino
	FindByID(int) []*Vino
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New ...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

func (s service) AddVino(m Vino) error {
	return nil
}

func (s service) FindByID(id int) []*Vino {
	var m []*Vino
	if err := s.db.Select(&m, "SELECT * FROM vinos where ID=$1", id); err != nil {
		panic(err)
	}
	return m
}

func (s service) FindAll() []*Vino {
	var list []*Vino
	if err := s.db.Select(&list, "SELECT * FROM vinos"); err != nil {
		panic(err)
	}
	return list
}

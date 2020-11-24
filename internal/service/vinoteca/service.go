package vinoteca

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/matiaspocai/sem-go-1/internal/config"
)

// Vino ...
type Vino struct {
	ID   int    `json:"ID"`
	Text string `json:"Text"`
}

// Service ...
type Service interface {
	FindAll() []*Vino
	FindByID(int) []*Vino
	DeleteVino(int) []*Vino
	PostVino(string) []*Vino
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New ...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

// FindByID busca por ide
func (s service) FindByID(id int) []*Vino {
	var m []*Vino
	if err := s.db.Select(&m, "SELECT * FROM vinos where ID=$1", id); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return m
}

// FindAll trae todos los vinos
func (s service) FindAll() []*Vino {
	var list []*Vino
	if err := s.db.Select(&list, "SELECT * FROM vinos"); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return list
}

// PostVino crea un vino
func (s service) PostVino(t string) []*Vino {
	var m []*Vino
	res := "INSERT INTO vinos (text) VALUES (?)"
	s.db.MustExec(res, t)
	return m
}

// DeleteVino elimina vino por id...
func (s service) DeleteVino(id int) []*Vino {
	var m []*Vino
	if err := s.db.Select(&m, "DELETE FROM vinos where ID=$1", id); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return m
}

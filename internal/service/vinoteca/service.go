package vinoteca

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/matiaspocai/sem-go-1/internal/config"
)

// Vino ...
type Vino struct {
	ID       int    `json:"ID"`
	Nombre   string `json:"Nombre"`
	Marca    string `json:"Marca"`
	Varietal string `json:"Varietal"`
	Precio   int    `json:"Precio"`
}

// Service ...
type Service interface {
	FindAll() []*Vino
	FindByID(int) []*Vino
	DeleteVino(int) []*Vino
	PostVino(Vino) []*Vino
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
	if err := s.db.Select(&m, "SELECT * FROM vinoteca where ID=$1", id); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return m
}

// FindAll trae todos los vinos
func (s service) FindAll() []*Vino {
	var list []*Vino
	if err := s.db.Select(&list, "SELECT * FROM vinoteca"); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return list
}

// PostVino crea un vino
func (s service) PostVino(v Vino) []*Vino {
	var mv []*Vino
	res := "INSERT INTO vinoteca (nombre, marca, varietal, precio) VALUES (?,?,?,?)"
	s.db.MustExec(res, v.Nombre, v.Marca, v.Varietal, v.Precio)
	return mv
}

// DeleteVino elimina vino por id...
func (s service) DeleteVino(id int) []*Vino {
	var m []*Vino
	if err := s.db.Select(&m, "DELETE FROM vinoteca where ID=$1", id); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return m
}

package vinoteca

import (
	"github.com/jmoiron/sqlx"
	"github.com/matiaspocai/sem-go-1/internal/config"
)

// Vino ...
type Vino struct {
	ID   int
	Text string
}

// Service ...
type Service interface {
	FindAll() []*Vino
	FindByID(int) []*Vino
	DeleteVino(int) []*Vino
	// PostVino(*gin.Context) error
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
		panic(err)
	}
	return m
}

// FindAll trae todos los vinos
func (s service) FindAll() []*Vino {
	var list []*Vino
	if err := s.db.Select(&list, "SELECT * FROM vinos"); err != nil {
		panic(err)
	}
	return list
}

/* // PostVino crea un vino
func (s service) PostVino(c *gin.Context) error {
	var v []*Vino
	c.BindJSON(&v)
	return v
} */

// DeleteVino elimina vino por id...
func (s service) DeleteVino(id int) []*Vino {
	var m []*Vino
	if err := s.db.Select(&m, "DELETE FROM vinos where ID=$1", id); err != nil {
		panic(err)
	}
	return m
}

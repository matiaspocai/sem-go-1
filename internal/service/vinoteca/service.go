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
	FindByID(int) *Vino
	DeleteVino(int) string
	PostVino(Vino) string
	PutVino(int, Vino) string
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New sqlx más configuración ...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

// FindAll retorna la Lista de productos, vinos
func (s service) FindAll() []*Vino {
	var list []*Vino
	if err := s.db.Select(&list, "SELECT * FROM vinoteca"); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return list
}

// FindByID busca y devuelve producto por identificador
func (s service) FindByID(id int) *Vino {
	var v Vino
	if err := s.db.Get(&v, "SELECT * FROM vinoteca where ID=$1", id); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return &v
}

// DeleteVino elimina vino por id y retorna un string...
func (s service) DeleteVino(id int) string {
	var str string
	if err := s.db.MustExec("DELETE FROM vinoteca where ID=$1", id); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return str
}

// PostVino crea producto (vino) y devuelve un mensaje
func (s service) PostVino(v Vino) string {
	res := "INSERT INTO vinoteca (nombre, marca, varietal, precio) VALUES (?,?,?,?)"
	s.db.MustExec(res, v.Nombre, v.Marca, v.Varietal, v.Precio)
	return "La inserción fue un éxito!"
}

// PutVino, edita productos por id y devuelve la lista actualizada
func (s service) PutVino(id int, v Vino) string {
	res := "UPDATE vinoteca SET nombre = ?, marca = ?, varietal = ?, precio = ? WHERE ID=id"
	s.db.MustExec(res, v.Nombre, v.Marca, v.Varietal, v.Precio)
	return "Producto editado con éxito."
}

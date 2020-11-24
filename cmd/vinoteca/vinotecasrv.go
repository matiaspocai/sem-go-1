package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/matiaspocai/sem-go-1/internal/config"
	"github.com/matiaspocai/sem-go-1/internal/database"
	"github.com/matiaspocai/sem-go-1/internal/service/vinoteca"
)

func main() {
	configFile := flag.String("config", "./config.yaml", "this is the service config")
	flag.Parse()

	cfg := config.LoadConfig(*configFile)

	db, err := database.NewDatabase(cfg)
	defer db.Close()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	service, _ := vinoteca.New(db, cfg)
	httpService := vinoteca.NewHTTPTransport(service)

	// createSchema(db)
	// eliminarTabla(db)

	r := gin.Default()
	httpService.Register(r)
	r.Run()
}

func eliminarTabla(db *sqlx.DB) error {
	res := `DROP TABLE vinoteca`
	db.MustExec(res)
	return nil
}

// agregar createSchema(db) para crear tabla/base de datos
func createSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS vinoteca (
		id integer primary key autoincrement,
		nombre varchar, 
		marca varchar,
		varietal varchar,
		precio int
		);`

	// execute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	// or, you can use MustExec, which panics on error
	res := `INSERT INTO vinoteca (nombre, marca, varietal, precio) VALUES (?,?,?,?)`
	n := "Elastic"
	m := "Style"
	v := "Malbec"
	p := 210
	db.MustExec(res, n, m, v, p)
	return nil
}

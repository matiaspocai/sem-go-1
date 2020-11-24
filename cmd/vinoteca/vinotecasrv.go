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

	r := gin.Default()
	httpService.Register(r)
	r.Run()
}

// agregar createSchema(db) para crear tabla/base de datos
func createSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS vinoteca (
		id integer primary key autoincrement,
		nombre varchar, 
		marca varchar,
		varietal varchar,
		precio int);`

	// execute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	// or, you can use MustExec, which panics on error
	insertMessage := `INSERT INTO vinoteca (nombre, marca, varietal, precio) VALUES (?,?,?,?)`
	n := "Libre"
	m := "Style"
	v := "Cabernet"
	p := 175
	db.MustExec(insertMessage, n, m, v, p)
	return nil
}

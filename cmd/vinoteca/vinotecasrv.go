package main

import (
	"flag"
	"fmt"
	"os"
	"time"

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

	// createSchema(db) : ejecutar para crear tabla por primera vez

	r := gin.Default()
	httpService.Register(r)
	r.Run()
}

// agregar createSchema(db) para crear tabla/base de datos
func createSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS vinos (
		id integer primary key autoincrement,
		text varchar);`

	// execute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	// or, you can use MustExec, which panics on error
	insertMessage := `INSERT INTO vinos (text) VALUES (?)`
	s := fmt.Sprintf("Message number %v", time.Now().Nanosecond())
	db.MustExec(insertMessage, s)
	return nil
}

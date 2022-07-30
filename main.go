package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/mekuanint12/simple_api/api"
	db "github.com/mekuanint12/simple_api/db/sqlc"
	"github.com/mekuanint12/simple_api/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cant load config file: ", err)
	}
	con, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Can not connect to the database:", err)
	}
	create_account := db.NewUser(con)
	server := api.NewServer(create_account)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Can not start server:", err)
	}

}

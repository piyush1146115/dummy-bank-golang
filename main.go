package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/piyush1146115/dummy-bank-golang/api"
	db "github.com/piyush1146115/dummy-bank-golang/db/sqlc"
	"github.com/piyush1146115/dummy-bank-golang/db/util"

	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can not connect to DB:", err)
	}
	if err := conn.Ping(); err != nil {
		log.Fatal("failed to ping db")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	if err := server.Start(config.SERVERAddress); err != nil {
		log.Fatal("Failed to start server")
	}

}

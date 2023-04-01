package main

import (
	"database/sql"
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
}

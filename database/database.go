package database

import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)

var DatabaseConnection *sql.DB

func ConnectDB() {
	var err error
	configurationString := "user = postgres dbname = Loja  password = 123 sslmode = disable"
	DatabaseConnection, err = sql.Open("postgres", configurationString)
	if err != nil {
		log.Println("Error: ", err)
	}

	DatabaseConnection.Ping()
	if err != nil {
		log.Println("Error: ", err)
	} else {
		log.Println("Conexão realizada com sucesso.")
	}
}








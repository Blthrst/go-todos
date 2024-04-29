package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var connectionString string

// setup funcs

func InitSecrets() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err.Error())
	}

	var dbUser string = os.Getenv("MYSQL_USER")

	var dbName string = os.Getenv("MYSQL_DBNAME")

	var dbPwd string = os.Getenv("MYSQL_PWD")

	var dbHost string = os.Getenv("MYSQL_HOST")

	var dbPort string = os.Getenv("MYSQL_PORT")

	if dbHost == "" {
		dbHost = "127.0.0.1"
	}

	if dbPort == "" {
		dbPort = "3306"
	}

	connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPwd, dbHost, dbPort, dbName)

}

func PrepareDatabase() {
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	_, err = db.Exec(`CREATE TABLE if not exists users (
		id int NOT NULL,
		name varchar(255) NOT NULL,
		PRIMARY KEY (id)
	  );`)

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Exec(`CREATE TABLE if not exists todos (
		id int NOT NULL,
		title varchar(255) NOT NULL,
		description longtext NOT NULL,
		is_done tinyint(1) NOT NULL,
		user_id int NOT NULL,
		PRIMARY KEY (id)
	  );`)

	if err != nil {
		log.Fatal(err.Error())
	}
}

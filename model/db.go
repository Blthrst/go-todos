package model

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var connectionString string

func InitSecrets() {
	err := godotenv.Load()

	if err != nil {
		panic(err.Error())
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

func GetUser(id int) (*User, error) {

	user := &User{}

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		return user, err
	}

	defer db.Close()

	rows, err := db.Query("select * from users where id = ?", id)

	if err != nil {
		return user, err
	}

	rows.Next()
	rows.Scan(&user.Id, &user.Name)

	return user, nil
}

func InsertUser(user User) (sql.Result, error) {
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	result, err := db.Exec("insert into users (id, name) values (?, ?);", user.Id, user.Name)

	if err != nil {
		return nil, err
	}

	return result, nil
}
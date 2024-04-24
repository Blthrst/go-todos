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

func GetUser(id int) (User, error) {

	user := &User{}

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		return *user, err
	}

	defer db.Close()

	rows, err := db.Query("select * from users where id = ?", id)

	if err != nil {
		return *user, err
	}

	rows.Next()
	rows.Scan(&user.Id, &user.Name)

	return *user, nil
}

func GetAllUsers() ([]User, error) {

	users := []User{}

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		return users, err
	}

	defer db.Close()

	rows, err := db.Query("select * from users")

	if err != nil {
		return users, err
	}

	for rows.Next() {
		user := &User{}

		rows.Scan(&user.Id, &user.Name)

		users = append(users[0:], *user)
	}

	return users, nil
}

func InsertUsers(users []User) (error) {
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		return err
	}

	defer db.Close()

	for i:=0; i < len(users); i++ {
		_, err := db.Exec("insert into users (id, name) values (?, ?);", users[i].Id, users[i].Name)

	if err != nil {
		return err
	}
	}

	return nil
}

func DeleteUsers(ids []int) error {
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		return err
	}

	defer db.Close()

	for i:=0; i < len(ids); i++ {
		_, err := db.Exec("delete from users where id = ?", ids[i])

		if err != nil {
			return err
		}
	}
	
	return nil
}

func UpdateUser(ub UpdateUserBody) error {
	
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("update users set id = ?, name = ? where id = ?", ub.User.Id, ub.User.Name, ub.Id)

	if err != nil {
		return err
	}
	
	return nil
}

func CreateTodo(todo Todo) error {
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("insert into todos (id, title, decription, is_done, user_id) values ({Id}, {Title}, {Description}, {IsDone}, {UserId})", todo)

	if err != nil {
		return err
	}

	return nil
}

func DeleteTodos(ids []int) error {

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		return err
	}

	defer db.Close()

	for i := 0; i < len(ids); i++ {
		_, err = db.Exec("delete from todos where id = ?", ids[i])

		if err != nil {
			return err
		}
	}

	return nil
}
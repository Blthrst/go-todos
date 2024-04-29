package model

import (
	"database/sql"
)

//users funcs

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

package model

import (
	"database/sql"
)

//todos funcs

func GetTodos() ([]Todo, error) {

	todos := []Todo{}

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("select * from todos")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		todo := &Todo{}

		err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.IsDone, &todo.UserId)

		if err != nil {
			return nil, err
		}

		todos = append(todos[0:], *todo)
	}

	return todos, nil
}

func CreateTodo(todo Todo) error {
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("insert into todos (id, title, description, is_done, user_id) values (?, ?, ?, ?, ?)", todo.Id, todo.Title, todo.Description, todo.IsDone, todo.UserId)

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

func CompleteTodo(id int) error {
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("update todos set is_done = true where id = ?", id)

	if err != nil {
		return err
	}

	return nil
}

func UpdateTodo(todo Todo) error {
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("update todos set title = ?, description = ?, is_done = ?, user_id = ? where id = ?", todo.Title, todo.Description, todo.IsDone, todo.UserId, todo.Id)

	if err != nil {
		return err
	}

	return nil
}
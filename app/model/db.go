package model

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// Todo data structure
type Todo struct {
	ID     int
	Action string
	Done   bool
}

// Db handle
var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("sqlite3", "./data/borgdir.media")
	if err != nil {
		panic(err)
	}
}

// GetAll Todo
func GetAll() (todos []Todo, err error) {
	rows, err := Db.Query("select id, action, done from todo")

	if err != nil {
		return
	}

	for rows.Next() {
		todo := Todo{}
		err = rows.Scan(&todo.ID, &todo.Action, &todo.Done)

		if err != nil {
			return
		}

		todos = append(todos, todo)
	}

	rows.Close()
	return
}

// Get Todo with the provided id
func Get(id int) (todo Todo, err error) {
	todo = Todo{}
	err = Db.QueryRow("select id, action, done from todo where id = $1", id).Scan(&todo.ID, &todo.Action, &todo.Done)
	return
}

// Add Todo
func (todo *Todo) Add() (err error) {
	statement := "insert into todo (action, done) values ($1, $2)"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(todo.Action, todo.Done)
	return
}

// ToggleStatus toggles the completion status of the Todo with the provided id
func (todo *Todo) ToggleStatus() (err error) {
	_, err = Db.Exec("update todo set done = not done where id = $1", todo.ID)
	return
}

// Delete Todo with the provided id from the list of Todos
func (todo *Todo) Delete() (err error) {
	_, err = Db.Exec("delete from todo where id = $1", todo.ID)
	return
}

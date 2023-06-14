package models

import (
	"log"
	"time"
)

type Todo struct {
	ID int
	Title string
	Content string
	UserID int
	CreatedAt time.Time
}

func (u *User) CreateTodo(title string, content string) (err error) {
	cmd := `insert into todos (
		content,
		title,
		user_id,
		created_at) values (?, ?, ?, ?)`
	_, err = Db.Exec(cmd, content, title, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetTodo(id int) (todo Todo, err error) {
	cmd := `select id, content, title, user_id, created_at from todos where id = ?`
	todo = Todo{}

	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.Title,
		&todo.UserID,
		&todo.CreatedAt)

	return todo, err
}

func GetTodos() (todos []Todo, err error) {
	cmd := `select id, content, title, user_id, created_at from todos`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.Title,
			&todo.UserID,
			&todo.CreatedAt)

		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()   //忘れない

	return todos, err
}

func (u *User) GetTodosByUser() (todos []Todo, titles []string, err error) {
	cmd := `select id, content, title, user_id, created_at from todos where user_id = ?`

	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.Title,
			&todo.UserID,
			&todo.CreatedAt)
			
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
		titles = append(titles, todo.Title)
	}
	rows.Close()
	return todos, titles, err
}

func (u *User) GetTodosTitleByUser() (titles []string, err error) {
	cmd := `select id, title, user_id, created_at from todos where user_id = ?`
	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.UserID,
			&todo.CreatedAt)

		if err != nil {
			log.Fatalln(err)
		}
		titles = append(titles, todo.Title)
	}
	defer rows.Close()
	return titles, err
}

func (t *Todo) UpdateTodo() error {
	cmd := `update todos set content = ?, title = ?, user_id = ? where id = ?`
	_, err = Db.Exec(cmd, t.Content, t.Title, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (t *Todo) DeleteTodo() error {
	cmd := `delete from todos where id = ?`
	_, err = Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
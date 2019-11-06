package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alexodorico/todo/db"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	_ "github.com/lib/pq"
)

type todo struct {
	ID   string `json:"id"`
	Item string `json:"item"`
}

func main() {
	db.InitDB()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/todos", func(r chi.Router) {
		r.Get("/", listTodos)
		r.Post("/", createTodo)

		r.Route("/{todoID}", func(r chi.Router) {
			r.Get("/", getTodo)
			r.Put("/", updateTodo)
			r.Delete("/", deleteTodo)
		})
	})

	http.ListenAndServe(":8080", r)
}

func listTodos(w http.ResponseWriter, r *http.Request) {
	var ts []todo

	rows, err := db.Conn.Query("SELECT * FROM todos")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var t todo

		err = rows.Scan(&t.ID, &t.Item)
		if err != nil {
			panic(err)
		}

		ts = append(ts, t)
	}

	render.JSON(w, r, ts)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var t todo
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	stmt := "INSERT INTO todos (item) VALUES ($1)"
	_, err = db.Conn.Exec(stmt, t.Item)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")

	var t todo
	row := db.Conn.QueryRow("SELECT * FROM todos WHERE id = $1", todoID)
	switch err := row.Scan(&t.ID, &t.Item); err {
	case sql.ErrNoRows:
		fmt.Println("No rows")
	case nil:
		render.JSON(w, r, t)
	default:
		panic(err)
	}
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	w.Write([]byte(todoID))
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	w.Write([]byte(todoID))
}

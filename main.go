package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

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
	w.Write([]byte("listTodos"))
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("createTodo"))
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	w.Write([]byte(todoID))
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	w.Write([]byte(todoID))
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	w.Write([]byte(todoID))
}

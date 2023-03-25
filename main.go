package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/HackStoga-2023-server/models"
	"github.com/HackStoga-2023-server/util"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "docker"
	dbname   = "postgres"
)

var DB_INSTANCE *sql.DB

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	DB_INSTANCE = db

	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Established a successful connection!")

	// Init Router
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/get_todo_items", GetItems).Methods("POST")
	r.HandleFunc("/api/v1/post_todo_items", AddItems).Methods("POST")
	r.HandleFunc("/api/v1/todo_items/{todo_id}", DeleteItem).Methods("DELETE")
	r.HandleFunc("/api/v1/login", SignIn).Methods("POST")
	r.HandleFunc("/api/v1/create_account", CreateAccount).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", r))
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	err := util.GetItems(w, r, DB_INSTANCE)
	if err != nil {
		var response = models.TodoItemJsonResponse{Type: "failure", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}
}

func AddItems(w http.ResponseWriter, r *http.Request) {
	err := util.AddItems(w, r, DB_INSTANCE)
	if err != nil {
		var response = models.TodoItemJsonResponse{Type: "failure", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	err := util.DeleteItem(w, r, DB_INSTANCE)
	if err != nil {
		var response = models.TodoItemJsonResponse{Type: "failure", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	err := util.Login(w, r, DB_INSTANCE)
	if err != nil {
		var response = models.TodoItemJsonResponse{Type: "failure", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	err := util.CreateAccount(w, r, DB_INSTANCE)
	if err != nil {
		var response = models.TodoItemJsonResponse{Type: "failure", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}
}

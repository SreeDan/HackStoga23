package util

import (
	"database/sql"
	"encoding/json"
	"github.com/HackStoga-2023-server/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func AddItems(w http.ResponseWriter, r *http.Request, db *sql.DB) error {
	// establishes context
	params := mux.Vars(r)
	ctx := r.Context()
	userID, _ := strconv.Atoi(params["id"])
	w.Header().Set("Content-Type", "application/json")

	//  Puts the body into a struct
	var itemStruct models.TodoItem
	err := json.NewDecoder(r.Body).Decode(&itemStruct)
	if err != nil {
		panic(err)
	}

	err = Authenticate(ctx, itemStruct.UserId, itemStruct.Username, itemStruct.Password, db)
	if err != nil {
		return err
	}

	insertStmt := `insert into todo_list(user_id, title, description) VALUES($1, $2, $3)`
	_, err = db.ExecContext(ctx, insertStmt, userID, itemStruct.Title, itemStruct.Description)
	if err != nil {
		return err
	}

	var response = models.TodoItemJsonResponse{Type: "success"}
	json.NewEncoder(w).Encode(response)

	return nil
}

func GetItems(w http.ResponseWriter, r *http.Request, db *sql.DB) error {
	// establishes context
	params := mux.Vars(r)
	ctx := r.Context()
	userID, _ := strconv.Atoi(params["id"])
	w.Header().Set("Content-Type", "application/json")

	var itemStruct models.TodoItem
	err := json.NewDecoder(r.Body).Decode(&itemStruct)
	if err != nil {
		panic(err)
	}

	err = Authenticate(ctx, itemStruct.UserId, itemStruct.Username, itemStruct.Password, db)
	if err != nil {
		return err
	}

	//  Makes sure db is online
	err = db.Ping()
	if err != nil {
		return err
	}

	var items []models.TodoItem
	sqlStmt := `SELECT * FROM todo_list where user_id = $1`
	rows, err := db.QueryContext(ctx, sqlStmt, userID)
	if err != nil {
		return err
	}

	for rows.Next() {
		var id int64
		var userId int64
		var title string
		var description string
		var dateCreated time.Time
		var dateUpdated time.Time
		var deadline *time.Time

		err = rows.Scan(&id, &userId, &title, &description, &dateCreated, &dateUpdated, &deadline)

		items = append(items, models.TodoItem{
			Id:          id,
			UserId:      userId,
			Title:       title,
			Description: description,
			DateCreated: dateCreated,
			DateUpdated: dateUpdated,
			Deadline:    deadline,
		})

	}

	result := []map[string]interface{}{}

	for _, v := range items {
		result = append(result, models.ItemSerialize(ctx, v))
	}

	var response = models.TodoItemJsonResponse{Type: "success", Data: items}
	json.NewEncoder(w).Encode(response)

	return nil
}

func DeleteItem(w http.ResponseWriter, r *http.Request, db *sql.DB) error {
	// establishes context
	params := mux.Vars(r)
	ctx := r.Context()
	todoId, _ := strconv.Atoi(params["todo_id"])
	w.Header().Set("Content-Type", "application/json")

	//  Puts the body into a struct
	var itemStruct models.TodoItem
	err := json.NewDecoder(r.Body).Decode(&itemStruct)
	if err != nil {
		panic(err)
	}

	err = Authenticate(ctx, itemStruct.UserId, itemStruct.Username, itemStruct.Password, db)
	if err != nil {
		return err
	}

	deleteStmt := `DELETE FROM todo_list WHERE user_id = $1 AND id = $2`
	_, err = db.ExecContext(ctx, deleteStmt, itemStruct.UserId, todoId)
	if err != nil {
		return err
	}

	var response = models.TodoItemJsonResponse{Type: "success"}
	json.NewEncoder(w).Encode(response)

	return nil
}

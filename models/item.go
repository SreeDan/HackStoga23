package models

import (
	"context"
	"time"
)

type TodoItemJsonResponse struct {
	Type    string     `json:"type"`
	Data    []TodoItem `json:"data"`
	Message string     `json:"message"`
}

type TodoItem struct {
	Id          int64      `json:"id" db:"id"`
	UserId      int64      `json:"user_id" db:"user_id"`
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description" db:"description"`
	DateCreated time.Time  `json:"date_created" db:"date_created"`
	DateUpdated time.Time  `json:"date_updated" db:"date_updated"`
	Deadline    *time.Time `json:"deadline" db:"deadline"`
	Username    string     `json:"username" db:"usernamed"`
	Password    string     `json:"password" db:"password"`
}

func ItemSerialize(ctx context.Context, ti TodoItem) map[string]interface{} {
	result := map[string]interface{}{
		"id":           ti.Id,
		"user_id":      ti.UserId,
		"title":        ti.UserId,
		"description":  ti.Description,
		"date_created": ti.DateCreated,
		"date_updated": ti.DateUpdated,
		"deadline":     ti.Deadline,
	}

	return result

}

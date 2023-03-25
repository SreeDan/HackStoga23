package models

import (
	"context"
)

type LoginResponse struct {
	Type    string     `json:"type"`
	Data    []TodoItem `json:"data"`
	Message string     `json:"message"`
}

type Login struct {
	UserId   int64  `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

func LoginSerialize(ctx context.Context, l Login) map[string]interface{} {
	result := map[string]interface{}{
		"user_id":  l.UserId,
		"username": l.Username,
		"password": l.Password,
	}

	return result

}

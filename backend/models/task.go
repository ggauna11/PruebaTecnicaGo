package models

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

// Task represents a task in the to-do list
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
type API struct {
	DB     *sql.DB
	Router *gin.Engine
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

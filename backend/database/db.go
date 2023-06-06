package database

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type API struct {
	DB     *sql.DB
	Router *gin.Engine
}

// InitDB initializes the database connection
func InitDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:Argentina.123@tcp(db:3306)/Tareas")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Println("Connected to the database")

	return db, nil
}

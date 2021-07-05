package main

import (
	/*
		"bytes"
		"fmt"
		"net/http"
	*/
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db, err := sql.Open("mysql", "root:yogi99yogi99@tcp(127.0.0.1:3306)/alearn_gin_golang")
	err = db.Ping()
	if err != nil {
		panic("Failed to connection on Database...")
	}
	defer db.Close()

	router := gin.Default()
	router.Run(":8080")
}
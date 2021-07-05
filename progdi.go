package main

import (
	/*
		"bytes"
		"fmt"
	*/

	"database/sql"
	"net/http"

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

	// modeling
	type Progdi struct {
		Id 				int 		`json: "id"`
		Jenjang 	string 	`json: "jenjang"`
		NmProgdi 	string 		`json: "nmprogdi"`
	}

	// show data by ID
	router.GET("/:id", func(c *gin.Context) {
		var (
			progdi Progdi
			result gin.H
		)

		id := c.Param("id")
		row := db.QueryRow("select id, jenjang, nmprogdi from progdi where id = ?;", id)
		err = row.Scan(&progdi.Id, &progdi.Jenjang, &progdi.NmProgdi)

		if err != nil {
			// if no result
			result = gin.H{
				"status" : "success",
				"message" : "no data",
			}
		} else {
			result = gin.H{
				"status" : "success",
				"message" : "found data",
				"data" : progdi,
			}
		}
		c.JSON(http.StatusOK, result)
	})

	// show all Data
	router.Run(":8080")
}
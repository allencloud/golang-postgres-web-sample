package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"os"
)

func main() {
	host := os.Getenv("POSTGRES_PORT_5432_TCP_ADDR")
	port := os.Getenv("POSTGRES_PORT_5432_TCP_PORT")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	dbname := os.Getenv("INSTANCE_NAME")

	connection_info := "host=" + host + " port=" + port + " user=" + username + " password=" + password + " dbname=" + dbname + " sslmode=disable"
	fmt.Println(connection_info)
	db, err := sql.Open("postgres", connection_info)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Exec("CREATE TABLE Persons(Name varchar(255))")

	err = db.Exec(fmt.Sprintf("INSERT INTO Persons VALUES (%s)", "DaoCloud"))

	rows, err := db.Query(`SELECT Name FROM Persons`)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, rows)
	})

	r.Run(":8080")
}

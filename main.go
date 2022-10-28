package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

type student struct {
	SlackUsername string `json:"slackUsername"`
	Backend       bool   `json:"backend"`
	Age           int    `json:"age"`
	Bio           string `json:"bio"`
}

var myDetails = student{SlackUsername: "eduonyia", Backend: true, Age: 35, Bio: "Ready to impact my world"}

func getStudent(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, myDetails)

}

func main() {
	router := gin.Default()

	router.GET("/student", getStudent) // first endpoint

	port := os.Getenv("PORT")
	if port == "" {
		port = "808"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}

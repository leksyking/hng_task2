package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Request struct {
	OperationType string `json:"operation_type"`
	X             int    `json:"x"`
	Y             int    `json:"y"`
}

var result int

func main() {
	server := gin.New()

	server.Use(gin.Logger(), gin.Recovery())

	server.POST("/", Operation)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)
}

func Operation(c *gin.Context) {
	var request Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "Something went wrong"})
		return
	}
	if request.OperationType == "addition" {
		result = request.X + request.Y
	} else if request.OperationType == "subtraction" {
		result = request.X - request.Y
	} else {
		result = request.X * request.Y
	}
	switch request.OperationType {
	case "addition":
		result = request.X + request.Y
	}
	response := gin.H{
		"slackUsername":  "leksyking",
		"operation_type": request.OperationType,
		"result":         result,
	}

	c.JSON(http.StatusOK, response)

}

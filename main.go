package main

import (
	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	router := gin.Default()

	// Define routes here

	router.Run()
}

package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todo struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var db *gorm.DB
var err error

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables.")
	}

	dsn := os.Getenv("DATABASE_URL")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&Todo{})

	router := gin.Default()

	router.GET("/todos", GetTodos)
	router.POST("/todos", CreateTodo)
	router.PUT("/todos/:id", UpdateTodo)
	router.DELETE("/todos/:id", DeleteTodo)

	router.Run()
}

func GetTodos(c *gin.Context) {
	var todos []Todo
	db.Find(&todos)
	c.JSON(200, todos)
}

func CreateTodo(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Create(&todo)
	c.JSON(201, todo)
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo Todo
	if err := db.First(&todo, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Save(&todo)
	c.JSON(200, todo)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	var todo Todo
	if err := db.First(&todo, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}
	db.Delete(&todo)
	c.JSON(200, gin.H{"message": "Todo deleted"})
}

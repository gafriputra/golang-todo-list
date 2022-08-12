package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gafriputra/todolist/activity"
	"github.com/gafriputra/todolist/handler"
	"github.com/gafriputra/todolist/todo"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// dsn := "gafri:gafri@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DBNAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	activityRepository := activity.NewRepository(db)
	todoRepository := todo.NewRepository(db)

	activityService := activity.NewService(activityRepository)
	todoService := todo.NewService(todoRepository)

	activityHandler := handler.NewActivityHandler(activityService)
	todoHandler := handler.NewTodoHandler(todoService)

	router := gin.Default()
	router.Use(cors.Default())
	apiActivity := router.Group("/activity-groups")
	apiTodo := router.Group("/todo-items")

	apiActivity.GET("/", activityHandler.GetActivities)
	apiActivity.GET("/:id", activityHandler.GetActivity)
	apiActivity.POST("/", activityHandler.CreateActivity)
	apiActivity.PATCH("/:id", activityHandler.UpdatedActivity)
	apiActivity.DELETE("/:id", activityHandler.DeleteActivity)

	apiTodo.GET("/", todoHandler.GetTodos)
	apiTodo.GET("/:id", todoHandler.GetTodo)
	apiTodo.POST("/", todoHandler.CreateTodo)
	apiTodo.PATCH("/:id", todoHandler.UpdatedTodo)
	apiTodo.DELETE("/:id", todoHandler.DeleteTodo)
}

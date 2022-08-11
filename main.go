package main

import (
	"log"

	"github.com/gafriputra/todolist/activity"
	"github.com/gafriputra/todolist/handler"
	"github.com/gafriputra/todolist/todo"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "gafri:gafri@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
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
	apiTodo.DELETE("/:id", todoHandler.UpdatedTodo)

	router.Run("0.0.0.0:3000")
}

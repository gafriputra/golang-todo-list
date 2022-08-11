package handler

import (
	"net/http"

	"github.com/gafriputra/todolist/helper"
	"github.com/gafriputra/todolist/todo"

	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	service todo.Service
}

func NewTodoHandler(service todo.Service) *todoHandler {
	return &todoHandler{service}
}

func (h *todoHandler) GetTodos(c *gin.Context) {
	Todos, err := h.service.GetTodos()
	if err != nil {
		response := helper.APIResponse("Error to get Todos", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of Todos", http.StatusOK, "success", todo.FormatTodos(Todos))
	c.JSON(http.StatusOK, response)
}

func (h *todoHandler) GetTodo(c *gin.Context) {
	var input todo.GetTodoDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of todo", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	todoDetail, err := h.service.GetTodoByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of todo", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if todoDetail.ID == 0 {
		response := helper.APIResponse("Todo Not Found", http.StatusNotFound, "success", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse("Success get todo detail", http.StatusOK, "success", todo.FormatTodo(todoDetail))
	c.JSON(http.StatusOK, response)

}

func (h *todoHandler) CreateTodo(c *gin.Context) {
	var input todo.CreateTodoInput
	response := helper.APIResponse

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		c.JSON(http.StatusBadRequest, response("Bad Request!", http.StatusUnprocessableEntity, "error", errorMessage))
		return
	}

	newTodo, err := h.service.CreateTodo(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, response("Failed Create Todo!", http.StatusBadRequest, "error", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response("Success Create Todo!", http.StatusOK, "success", todo.FormatTodo(newTodo)))
}

func (h *todoHandler) UpdatedTodo(c *gin.Context) {
	var inputID todo.GetTodoDetailInput
	response := helper.APIResponse

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update todo", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData todo.CreateTodoInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update todo", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedTodo, err := h.service.UpdateTodo(inputID, inputData)
	if err != nil {
		c.JSON(http.StatusBadRequest, response("Failed Update Todo!", http.StatusBadRequest, "error", err.Error()))
		return
	}
	c.JSON(http.StatusOK, response("Success Update Todo!", http.StatusOK, "success", todo.FormatTodo(updatedTodo)))
}

func (h *todoHandler) DeleteTodo(c *gin.Context) {
	var inputID todo.GetTodoDetailInput
	response := helper.APIResponse

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update todo", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData todo.CreateTodoInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update todo", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedTodo, err := h.service.UpdateTodo(inputID, inputData)
	if err != nil {
		c.JSON(http.StatusBadRequest, response("Failed Update Todo!", http.StatusBadRequest, "error", err.Error()))
		return
	}
	c.JSON(http.StatusOK, response("Success Update Todo!", http.StatusOK, "success", todo.FormatTodo(updatedTodo)))
}

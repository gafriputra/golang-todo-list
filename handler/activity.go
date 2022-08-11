package handler

import (
	"net/http"

	"github.com/gafriputra/todolist/activity"
	"github.com/gafriputra/todolist/helper"

	"github.com/gin-gonic/gin"
)

type activityHandler struct {
	service activity.Service
}

func NewActivityHandler(service activity.Service) *activityHandler {
	return &activityHandler{service}
}

func (h *activityHandler) GetActivities(c *gin.Context) {
	Activities, err := h.service.GetActivities()
	if err != nil {
		response := helper.APIResponse("Error to get Activities", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of Activities", http.StatusOK, "success", activity.FormatActivities(Activities))
	c.JSON(http.StatusOK, response)
}

func (h *activityHandler) GetActivity(c *gin.Context) {
	var input activity.GetActivityDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of activity", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	activityDetail, err := h.service.GetActivityByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of activity", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if activityDetail.ID == 0 {
		response := helper.APIResponse("Activity Not Found", http.StatusNotFound, "success", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse("Success get activity detail", http.StatusOK, "success", activity.FormatActivity(activityDetail))
	c.JSON(http.StatusOK, response)

}

func (h *activityHandler) CreateActivity(c *gin.Context) {
	var input activity.CreateActivityInput
	response := helper.APIResponse

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		c.JSON(http.StatusBadRequest, response("Bad Request!", http.StatusUnprocessableEntity, "error", errorMessage))
		return
	}

	newActivity, err := h.service.CreateActivity(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, response("Failed Create Activity!", http.StatusBadRequest, "error", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response("Success Create Activity!", http.StatusOK, "success", activity.FormatActivity(newActivity)))
}

func (h *activityHandler) UpdatedActivity(c *gin.Context) {
	var inputID activity.GetActivityDetailInput
	response := helper.APIResponse

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update activity", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData activity.CreateActivityInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update activity", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedActivity, err := h.service.UpdateActivity(inputID, inputData)
	if err != nil {
		c.JSON(http.StatusBadRequest, response("Failed Update Activity!", http.StatusBadRequest, "error", err.Error()))
		return
	}
	c.JSON(http.StatusOK, response("Success Update Activity!", http.StatusOK, "success", activity.FormatActivity(updatedActivity)))
}

func (h *activityHandler) DeleteActivity(c *gin.Context) {
	var inputID activity.GetActivityDetailInput
	response := helper.APIResponse

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update activity", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData activity.CreateActivityInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update activity", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedActivity, err := h.service.UpdateActivity(inputID, inputData)
	if err != nil {
		c.JSON(http.StatusBadRequest, response("Failed Update Activity!", http.StatusBadRequest, "error", err.Error()))
		return
	}
	c.JSON(http.StatusOK, response("Success Update Activity!", http.StatusOK, "success", activity.FormatActivity(updatedActivity)))
}

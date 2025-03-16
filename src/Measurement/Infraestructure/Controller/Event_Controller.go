package controller

import (
	application "NoisEsub/src/Measurement/Application"
	entities "NoisEsub/src/Measurement/Domain/Entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebhookController struct {
	service *application.EventServices
}

func NewWebhookController(h *application.EventServices) *WebhookController {
	return &WebhookController{service: h}
}


func (h *WebhookController) controller(c *gin.Context) {
	 var event entities.Event
	 if err := c.ShouldBindJSON(&event); err != nil{
		c.JSON((http.StatusBadRequest), gin.H{"error": err.Error()})
		return
	 }

	 if err := h.service.ProcessEvent(event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	 }
	 c.JSON(http.StatusOK, gin.H{"message": "Event received and processed"})
}


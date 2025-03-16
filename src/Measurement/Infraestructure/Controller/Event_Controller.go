package controller

import (
	application "NoisEsub/src/Measurement/Application"
	entities "NoisEsub/src/Measurement/Domain/Entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebhookController struct {
	service *application.EventService
}

func NewWebhookController(service *application.EventService) *WebhookController {
	return &WebhookController{service: service}
}

func (c *WebhookController) Handle(ctx *gin.Context) {
	var event entities.Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	if err := c.service.ProcessEvent(event); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event processed successfully"})
}
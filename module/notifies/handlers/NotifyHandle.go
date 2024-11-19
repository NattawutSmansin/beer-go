package handlers

import (
	"beer/module/notifies/models"
	"beer/module/notifies/usecases"
	"time"

	"github.com/gin-gonic/gin"
)

type NotifyHandler interface {
	Store(content gin.H) error
}

type notifyDataHandler struct {
	notifyUsecase usecases.NotifyUsecase
}

func NewNotifyHttpHandler(notifyUsecase usecases.NotifyUsecase) NotifyHandler {
	return &notifyDataHandler{
		notifyUsecase: notifyUsecase,
	}
}

func (n *notifyDataHandler) Store(content gin.H) error {
	Data := &models.CreateNotifyGo{
		Title:     content["Title"].(string),
		Detail:    content["Detail"].(string),
		CreatedAt: time.Now(),
	}

	// Set creation time
	Data.CreatedAt = time.Now()

	// Process the notification data
	err := n.notifyUsecase.NotifyDataProcessing(Data)
	if err != nil {
		return err
	}

	return nil
}

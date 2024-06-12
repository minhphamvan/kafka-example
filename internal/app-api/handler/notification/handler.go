package notificationhandler

import (
	"kafka-example/internal/service-notification"
)

type Handler struct {
	notificationService servicenotification.IService
}

const traceTag = "[notificationHandler]"

func New(
	notificationService servicenotification.IService,
) *Handler {
	return &Handler{
		notificationService: notificationService,
	}
}

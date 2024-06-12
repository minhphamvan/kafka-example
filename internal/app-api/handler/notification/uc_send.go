package notificationhandler

import (
	"fmt"
	"net/http"

	"kafka-example/internal/service-notification/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Send() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Parse request
		req := request.SendNotificationReq{}
		err := ctx.BindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("failed to bind JSON: %s", err.Error()),
			})
			return
		}

		// Send notification
		err = h.notificationService.SendNotification(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Notification sent successfully!",
		})
	}
}

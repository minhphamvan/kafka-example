package notificationhandler

import (
	"net/http"
	"strconv"

	"kafka-example/internal/service-notification/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetByUserID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Parse request
		userIDStr := ctx.Param("user_id")
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid user_id",
			})
			return
		}

		// Get notifications
		notifications, err := h.notificationService.GetByUserID(
			ctx, request.GetByUserIDReq{
				UserID: userID,
			},
		)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"notifications": notifications,
			"count":         len(notifications),
		})
	}
}

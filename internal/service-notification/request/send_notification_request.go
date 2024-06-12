package request

type SendNotificationReq struct {
	FromUserID int    `json:"from_user_id"`
	ToUserID   int    `json:"to_user_id"`
	Message    string `json:"message"`
}

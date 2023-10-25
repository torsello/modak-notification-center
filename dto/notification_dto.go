package dto

type NotificationDto struct {
	Type string `json:"type"`
	Receiver string `json:"receiver"`
	Message string `json:"message"` 
}

type NotificationsDto []NotificationDto

type NotificationResponseDto struct {
	Type string `json:"type"`
	Receiver string `json:"receiver"`
	Message string `json:"message"` 
	Status string `json:"status"` 
}

type NotificationsResponseDto []NotificationResponseDto

type DataResponseNotifications struct {
	Notifications NotificationsResponseDto `json:"notifications"`
}

type DataResponse struct {
	Data DataResponseNotifications `json:"data"`
}

type Data struct {
	Notifications NotificationsDto `json:"notifications"`
}

type DataRequest struct {
	Data Data `json:"data"`
}
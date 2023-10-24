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
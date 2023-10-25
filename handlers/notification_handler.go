package handlers

import (
	"encoding/json"
	"modak-notification-center/dto"
	"net/http"

	"modak-notification-center/services"
)


func SendNotification(response http.ResponseWriter, request *http.Request)  {
	response.Header().Set("Content-Type", "application/json")

	var notificationsData dto.DataRequest

	if err := json.NewDecoder(request.Body).Decode(&notificationsData); err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(map[string]interface{}{
				"data": map[string]string{
					"status":    "error",
					"code":      "000.000.000",
					"exception": "unnexpected_issue",
				},
			})
		return
	}

	notifications := notificationsData.Data.Notifications

	for _, notification := range notifications {
		if notification.Type == "" || notification.Receiver == "" || notification.Message == "" {
			response.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(response).Encode(map[string]interface{}{
				"data": map[string]string{
					"status": "error",
					"code": "000.000.001",
					"exception": "required_field_is_missing",
				},
			})
			return
		}
	}

	var responseDto dto.DataResponse
	service := services.NotificationServiceImpl{Gateway: services.Gateway{}}
	
	for _, notification := range notifications {
		status := "failed"
		ok := service.Send(notification.Type, notification.Receiver, notification.Message)
		
		if ok {
			status = "successful"
		}

		responseDto.Data.Notifications = append(responseDto.Data.Notifications, dto.NotificationResponseDto{
			Type: notification.Type,
			Receiver: notification.Receiver,
			Message: notification.Message,
			Status: status,
		})
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(responseDto)
	
}

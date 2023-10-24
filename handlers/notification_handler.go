package handlers

import (
	"encoding/json"
	"net/http"
	"rate-limited-notification/dto"

	"rate-limited-notification/services"
)


func SendNotification(response http.ResponseWriter, request *http.Request)  {
	response.Header().Set("Content-Type", "application/json")
	var notifications dto.NotificationsDto

	if err := json.NewDecoder(request.Body).Decode(&notifications); err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(map[string]string{
			"status":"error",
			"code": "000.000.000",
			"exception": "unnexpected_issue",
		})
		return
	}

	for _, notification := range notifications {
		if notification.Type == "" || notification.Receiver == "" || notification.Message == "" {
			response.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(response).Encode(map[string]string{
				"status": "error",
				"code": "000.000.001",
				"exception": "required_field_is_missing",
			})
			return
		}
	}

	var responseDto dto.NotificationsResponseDto
	service := services.NotificationServiceImpl{Gateway: services.Gateway{}}
	
	for _, notification := range notifications {
		status := "failed"
		ok := service.Send(notification.Type, notification.Receiver, notification.Message)
		
		if ok {
			status = "successful"
		}

		responseDto = append(responseDto, dto.NotificationResponseDto{
			Type: notification.Type,
			Receiver: notification.Receiver,
			Message: notification.Message,
			Status: status,
		})
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(responseDto)
	
}

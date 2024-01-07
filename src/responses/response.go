package responses

import "github.com/devkishor8007/email-sender/src/models"

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type SuccessResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    models.EmailSender `json:"data"`
}

type SuccessResponseList struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    []models.EmailSender `json:"data"`
}

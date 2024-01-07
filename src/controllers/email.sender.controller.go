package controllers

import (
	"context"
	"time"

	"github.com/devkishor8007/email-sender/src/database"
	"github.com/devkishor8007/email-sender/src/models"
	"github.com/devkishor8007/email-sender/src/responses"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateEmailTemplates(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var inputs models.EmailSender
	defer cancel()

	if err := c.Bind(&inputs); err != nil {
		return c.JSON(400, responses.ErrorResponse{Status: 400, Message: "Bad Request", Data: err.Error()})
	}

	newEmailTemplate := models.EmailSender{
		ID:         primitive.NewObjectID(),
		Subject:    inputs.Subject,
		Body:       inputs.Body,
		Created_At: primitive.NewDateTimeFromTime(time.Now()),
	}

	_, err := database.GetCollection("email_templates").InsertOne(ctx, newEmailTemplate)

	if err != nil {
		return c.JSON(400, responses.ErrorResponse{Status: 500, Message: "err", Data: err.Error()})
	}

	return c.JSON(201, responses.SuccessResponse{Status: 201, Message: "created successfully", Data: newEmailTemplate})
}

func GetEmailTemplates(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var emailSender []models.EmailSender
	defer cancel()

	filter := bson.M{}

	sort := bson.M{
		"created_at": -1,
	}

	pipeline := []bson.M{
		{"$match": filter},
		{"$sort": sort},
	}
	result, err := database.GetCollection("email_templates").Aggregate(ctx, pipeline)

	for result.Next(ctx) {
		var singleEmailResult models.EmailSender
		if err = result.Decode(&singleEmailResult); err != nil {
			return c.JSON(500, responses.ErrorResponse{Status: 500, Message: "error", Data: err.Error()})
		}

		emailSender = append(emailSender, singleEmailResult)
	}
	return c.JSON(201, responses.SuccessResponseList{Status: 200, Message: "success", Data: emailSender})
}

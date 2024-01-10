package controllers

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/devkishor8007/email-sender/src/database"
	"github.com/devkishor8007/email-sender/src/models"
	"github.com/devkishor8007/email-sender/src/responses"
	"github.com/devkishor8007/email-sender/src/utilis"
	"github.com/devkishor8007/email-sender/src/utilis/validate"
	"github.com/go-co-op/gocron/v2"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RunJob() {
	s, _ := gocron.NewScheduler()

	j, _ := s.NewJob(
		gocron.DurationJob(2*time.Second),
		gocron.NewTask(
			func(a string, b int) {
				fmt.Println(a, b)
			},
			"hello",
			1,
		),
	)

	fmt.Println(j.ID())

	fmt.Println("start....")
	go s.Start()
}

func TestJob(c echo.Context) error {
	RunJob()

	return c.JSON(201, responses.SuccessResponse{Status: 201, Message: "test job"})

}

func CreateEmailTemplates(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var inputs models.EmailTemplate
	defer cancel()

	if err := c.Bind(&inputs); err != nil {
		return c.JSON(400, responses.ErrorResponse{Status: 400, Message: "Bad Request", Data: err.Error()})
	}

	validateStatus := validate.ValidateStatus(inputs.Status)

	if validateStatus != nil {
		return c.JSON(400, responses.ErrorResponse{Status: 400, Message: "Bad Request", Data: validateStatus.Error()})
	}

	newEmailTemplate := models.EmailTemplate{
		ID:         primitive.NewObjectID(),
		Topic:      inputs.Topic,
		Status:     inputs.Status,
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var emailTemplates []models.EmailTemplate
	defer cancel()

	domain := os.Getenv("DOMAIN_URL")
	mailgun_api_key := os.Getenv("MAILGUN_API_KEY")
	from_email := os.Getenv("FROM_EMAIL")

	if len(domain) < 0 || len(mailgun_api_key) < 0 || len(from_email) < 0 {
		log.Fatal("empty domain or mailgun api key")
	}

	emailResponse := utilis.SendSimpleMessageP(domain, mailgun_api_key, from_email)

	fmt.Println(emailResponse)

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
		var singleEmailTemplate models.EmailTemplate
		if err = result.Decode(&singleEmailTemplate); err != nil {
			return c.JSON(500, responses.ErrorResponse{Status: 500, Message: "error", Data: err.Error()})
		}

		emailTemplates = append(emailTemplates, singleEmailTemplate)
	}
	return c.JSON(201, responses.SuccessResponseList{Status: 200, Message: "success", Data: emailTemplates})
}

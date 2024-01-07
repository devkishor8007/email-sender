package models

import (
	"github.com/devkishor8007/email-sender/src/utilis/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmailTemplate struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Topic      string             `json:"topic" bson:"topic"`
	Status     enums.Status       `json:"status" bson:"status"`
	Subject    string             `json:"subject" bson:"subject"`
	Body       string             `json:"body" bson:"body"`
	Created_At primitive.DateTime `json:"created_at"`
}

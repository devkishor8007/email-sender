package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type EmailSender struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Subject    string             `json:"subject" bson:"subject"`
	Body       string             `json:"body" bson:"body"`
	Created_At primitive.DateTime `json:"created_at"`
}

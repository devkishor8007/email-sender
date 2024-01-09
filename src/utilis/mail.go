package utilis

import (
	"context"
	"fmt"
	"log"

	"time"

	"github.com/mailgun/mailgun-go/v4"
)

func SendSimpleMessageP(domain, apiKey, from_email string) string {
	mg := mailgun.NewMailgun(domain, apiKey)
	fmt.Println("mess >>", mg)
	m := mg.NewMessage(
		from_email,
		"Hello kishor kc",
		"Testing some Mailgun new one!",
		"sender-semail@gmail.com",
	)

	m.SetRequireTLS(true)
	m.SetSkipVerification(true)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	message, _, err := mg.Send(ctx, m)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(message)

	return "email has been sent successfully!!"
}

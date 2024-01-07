package main

import (
	"github.com/devkishor8007/email-sender/src/routes"
)

func main() {
	e := routes.Route()

	e.Logger.Fatal(e.Start(":4000"))
}

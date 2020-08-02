package main

import (
	"os"

	app "github.com/mahesadhanaa/golang-restful/app"
)

func main() {
	a := app.App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8000")
}

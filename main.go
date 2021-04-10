// main.go

package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		APP_DB_USERNAME = os.Getenv("APP_DB_USERNAME")
		APP_DB_PASSWORD = os.Getenv("APP_DB_PASSWORD")
		APP_DB_NAME     = os.Getenv("APP_DB_NAME")
	)

	fmt.Printf("APP_DB_USERNAME: %s\nAPP_DB_PASSWORD: %s\nAPP_DB_NAME: %s\n", APP_DB_USERNAME, APP_DB_PASSWORD, APP_DB_NAME)

	a := App{}

	a.Initialize(
		APP_DB_USERNAME,
		APP_DB_PASSWORD,
		APP_DB_NAME)
	a.Run(":8010")
}

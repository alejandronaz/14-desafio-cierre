package main

import (
	"app/app/internal/application"
	"fmt"
	"os"
)

func main() {

	// export SERVER_ADDR="8080" && export DB_FILE="app/data/tickets.csv" && go run app/cmd/server/main.go

	// application
	// - config
	cfg := &application.ConfigAppDefault{
		ServerAddr: os.Getenv("SERVER_ADDR"),
		DbFile:     os.Getenv("DB_FILE"),
	}
	app := application.NewApplicationDefault(cfg)

	// - setup
	err := app.SetUp()
	if err != nil {
		fmt.Println(err)
		return
	}

	// - run
	err = app.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}

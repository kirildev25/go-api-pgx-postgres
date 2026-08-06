package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/kirildev25/go-api-pgx-postgres/internal/app"
	"github.com/kirildev25/go-api-pgx-postgres/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "enter port value without colon")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Logger.Printf("Running on port %d\n", port)

	r := routes.SetupRoutes(app)

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: r,
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		app.Logger.Fatal("can't serve")
	}
}

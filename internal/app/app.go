package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kirildev25/go-api-pgx-postgres/internal/api"
	"github.com/kirildev25/go-api-pgx-postgres/internal/store"
	"github.com/kirildev25/go-api-pgx-postgres/migrations"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

func NewApplication() (*Application, error) {
	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// stores
	workoutStore := store.NewPostgresWorkoutStore(pgDB)

	// handlers
	workoutHandler := api.NewWorkoutHandler(workoutStore)

	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		DB:             pgDB,
	}

	return app, nil
}

func (a Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available\n")
}

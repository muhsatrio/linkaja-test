package initiator

import (
	"fmt"
	"os"

	"financial-planner-be/internal/glue/routing"
	"financial-planner-be/internal/handler/rest"
	"financial-planner-be/internal/module/user"
	"financial-planner-be/internal/repository"
	"financial-planner-be/internal/storage/persistence"
	"financial-planner-be/platform/routers"

	"github.com/iDevoid/cptx"
	"github.com/sirupsen/logrus"
)

const (
	postgresURL = "postgresql://%s:%s@%s/iDevoid-db?sslmode=disable"

	domain = "user"
)

// User initializes the domain user
func User(testInit bool) {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbURL := fmt.Sprintf(postgresURL, dbUser, dbPass, dbHost)
	psql := cptx.Initialize(dbURL, dbURL, domain)
	postgresDB, postgresTX := psql.Open()

	databaseUser := persistence.UserInit(postgresDB)
	databaseProfile := persistence.ProfileInit(postgresDB)

	encryptKey := os.Getenv("ENCRYPTION_KEY")
	repo := repository.UserInit(encryptKey)

	usecase := user.Initialize(postgresTX, repo, databaseUser, databaseProfile)

	handler := rest.UserInit(usecase)
	router := routing.UserRouting(handler)

	port := os.Getenv("HOST_PORT")
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	server := routers.Initialize(port, allowedOrigins, router, domain)

	if testInit {
		logrus.Info("Initialize test mode Finished!")
		os.Exit(0)
	}

	server.Serve()
}

package main

import (
	"context"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/matyukhin00/pvz_service/internal/app"
	userRepo "github.com/matyukhin00/pvz_service/internal/repository/user"
	userService "github.com/matyukhin00/pvz_service/internal/service/user"
	"github.com/sirupsen/logrus"
)

var dbDNS string

func main() {
	ctx := context.Background()

	logger := logrus.New()

	pool, err := pgxpool.Connect(ctx, dbDNS)
	if err != nil {
		logger.Fatalf("Failed to connect database with DNS: `%s`", dbDNS)
	}
	defer pool.Close()

	repo := userRepo.NewUserRepository(pool)
	service := userService.NewUserService(repo)

	s := app.NewServer(logger, service)

	http.ListenAndServe(":8080", s)

}

func init() {
	godotenv.Load(".env")

	dbDNS = os.Getenv("PG_DNS")
}

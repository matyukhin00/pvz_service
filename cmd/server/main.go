package main

import (
	"context"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/matyukhin00/pvz_service/internal/app"
	pvzRepo "github.com/matyukhin00/pvz_service/internal/repository/pvz"
	userRepo "github.com/matyukhin00/pvz_service/internal/repository/user"
	pvzService "github.com/matyukhin00/pvz_service/internal/service/pvz"
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

	userRepo := userRepo.NewUserRepository(pool)
	userService := userService.NewUserService(userRepo)

	pvzRepo := pvzRepo.NewPvzRepository(pool)
	pvzService := pvzService.NewPvzService(pvzRepo)

	s := app.NewServer(logger, userService, pvzService)

	http.ListenAndServe(":8080", s)

}

func init() {
	godotenv.Load(".env")

	dbDNS = os.Getenv("PG_DNS")
}

package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matyukhin00/pvz_service/internal/service"
	"github.com/sirupsen/logrus"
)

type server struct {
	router      *mux.Router
	logger      *logrus.Logger
	userService service.UserService
	pvzService  service.PvzService
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/dummyLogin", s.handleDummyLogin()).Methods("POST")
	s.router.HandleFunc("/register", s.handleRegister()).Methods("POST")
	s.router.HandleFunc("/login", s.handleLogin()).Methods("POST")

	authRouter := s.router.PathPrefix("/").Subrouter()
	authRouter.Use(s.CheckJWT)
	authRouter.HandleFunc("/pvz", s.handlePvz()).Methods("POST")
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer(logger *logrus.Logger, userService service.UserService, pvzService service.PvzService) *server {
	s := &server{
		router:      mux.NewRouter(),
		logger:      logger,
		userService: userService,
		pvzService:  pvzService,
	}

	s.configureRouter()

	return s
}

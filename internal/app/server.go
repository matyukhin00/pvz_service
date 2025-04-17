package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matyukhin00/pvz_service/internal/service"
	"github.com/sirupsen/logrus"
)

type server struct {
	router           *mux.Router
	logger           *logrus.Logger
	userService      service.UserService
	pvzService       service.PvzService
	receptionService service.ReceptionService
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/dummyLogin", s.handleDummyLogin()).Methods("POST")
	s.router.HandleFunc("/register", s.handleRegister()).Methods("POST")
	s.router.HandleFunc("/login", s.handleLogin()).Methods("POST")

	authRouter := s.router.PathPrefix("/").Subrouter()
	authRouter.Use(s.CheckJWT)
	authRouter.HandleFunc("/pvz", s.handlePvz()).Methods("POST")
	authRouter.HandleFunc("/receptions", s.handleReceptions()).Methods("POST")
	authRouter.HandleFunc("/receptions/{pvzId}/close_last_reception", s.handleCloseLastReception()).Methods("POST")
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer(
	logger *logrus.Logger,
	userService service.UserService,
	pvzService service.PvzService,
	receptionService service.ReceptionService,
) *server {
	s := &server{
		router:           mux.NewRouter(),
		logger:           logger,
		userService:      userService,
		pvzService:       pvzService,
		receptionService: receptionService,
	}

	s.configureRouter()

	return s
}

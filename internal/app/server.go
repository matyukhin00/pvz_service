package app

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/matyukhin00/pvz_service/internal/service"
	"github.com/sirupsen/logrus"
)

var secretKey string

type server struct {
	router      *mux.Router
	logger      *logrus.Logger
	userService service.UserService
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/dummyLogin", s.handleDummyLogin()).Methods("POST")
	s.router.HandleFunc("/register", s.handleRegister()).Methods("POST")
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer(logger *logrus.Logger, userService service.UserService) *server {
	s := &server{
		router:      mux.NewRouter(),
		logger:      logger,
		userService: userService,
	}

	s.configureRouter()

	return s
}

func init() {
	secretKey = os.Getenv("secretKey")
}

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
	productService   service.ProductService
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/dummyLogin", s.handleDummyLogin()).Methods("POST")
	s.router.HandleFunc("/register", s.handleRegister()).Methods("POST")
	s.router.HandleFunc("/login", s.handleLogin()).Methods("POST")

	authRouter := s.router.PathPrefix("/").Subrouter()
	authRouter.Use(s.CheckJWT)
	authRouter.HandleFunc("/pvz", s.handlePvz()).Methods("POST")
	authRouter.HandleFunc("/receptions", s.handleReceptions()).Methods("POST")
	authRouter.HandleFunc("/pvz/{pvzId}/close_last_reception", s.handleCloseLastReception()).Methods("POST")
	authRouter.HandleFunc("/products", s.handleProducts()).Methods("POST")
	authRouter.HandleFunc("/pvz/{pvzId}/delete_last_product", s.handleDeleteLastProduct()).Methods("POST")
	authRouter.HandleFunc("/pvz", s.handlePvzPagination()).Methods("GET")
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) Run() {
	ch := make(chan error)
	go func() {
		ch <- http.ListenAndServe(":8080", s)
	}()
	s.logger.Infof("Server is running on localhost:8080")
	<-ch
}

func NewServer(
	logger *logrus.Logger,
	userService service.UserService,
	pvzService service.PvzService,
	receptionService service.ReceptionService,
	productService service.ProductService,
) *server {
	s := &server{
		router:           mux.NewRouter(),
		logger:           logger,
		userService:      userService,
		pvzService:       pvzService,
		receptionService: receptionService,
		productService:   productService,
	}

	s.configureRouter()

	return s
}

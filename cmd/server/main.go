package main

import (
	"net/http"

	"github.com/matyukhin00/pvz_service/internal/app"
)

func main() {
	s := app.NewServer()

	http.ListenAndServe(":8080", s)

}

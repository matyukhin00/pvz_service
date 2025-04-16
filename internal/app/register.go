package app

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/matyukhin00/pvz_service/internal/model"
)

func (s *server) handleRegister() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*3)
		defer cancel()

		req := model.User{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.Error{Message: "Invalid json format"})
			return
		}

		ans, err := s.userService.Create(ctx, req)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.Error{Message: err.Error()})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(model.RegisteredUser{
			Id:    ans.Id,
			Email: ans.Email,
			Role:  ans.Role,
		})
	}
}

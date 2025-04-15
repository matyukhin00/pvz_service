package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/matyukhin00/pvz_service/internal/model"
	"github.com/matyukhin00/pvz_service/internal/utils"
)

func (s *server) handleDummyLogin() http.HandlerFunc {
	type request struct {
		Role string `json:"role"`
	}
	type answer400 struct {
		Message string `json:"message"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(answer400{Message: "Content-Type must be application/json"})
			return
		}

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(answer400{Message: "Incorrect request"})
			return
		}

		if req.Role != "employee" && req.Role != "moderator" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(answer400{Message: "Role must be 'employee' or 'moderator'"})
			return
		}

		token, err := utils.GenerateToken(
			model.UserClaims{
				Role: req.Role,
			},
			[]byte(secretKey),
			time.Hour*24,
		)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(answer400{Message: "Failed to generate token"})
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(token))
		w.Header().Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}
}

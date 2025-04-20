package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/matyukhin00/pvz_service/internal/model"
)

// @Summary      Получение тестового токена
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body model.DummyLogin true "Роль пользователя"
// @Success      200 {string} string "Bearer token"
// @Failure      400 {object} model.Error "Некорректный запрос или роль"
// @Failure      500 {object} model.Error "Ошибка генерации токена"
// @Router 		 /dummyLogin [post]
func (s *server) handleDummyLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.Error{Message: "Content-Type must be application/json"})
			return
		}

		req := &model.DummyLogin{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.Error{Message: "Incorrect request"})
			return
		}

		if req.Role != "employee" && req.Role != "moderator" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.Error{Message: "Role must be 'employee' or 'moderator'"})
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*3)
		defer cancel()

		token, err := s.userService.DummyLogin(ctx, model.UserClaims{
			Role: req.Role,
		})

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(model.Error{Message: "Failed to generate token"})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Add("Authorization", fmt.Sprintf("Bearer %s", token))
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(token)

	}
}

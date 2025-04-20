package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/matyukhin00/pvz_service/internal/model"
)

// @Summary      Авторизация пользователя
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body model.Login true "Роль пользователя"
// @Success      200 {string} string "Bearer token"
// @Failure      400 {object} model.Error "Невалидный JSON"
// @Failure      401 {object} model.Error "Неверные данные для входа"
// @Router 		 /login [post]“
func (s *server) handleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		login := model.User{}
		err := json.NewDecoder(r.Body).Decode(&login)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.Error{Message: "Invalid json format"})
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*3)
		defer cancel()

		token, err := s.userService.Login(ctx, login)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.Error{Message: err.Error()})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Add("Authorization", fmt.Sprintf("Bearer %s", token))
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(token)
	}
}

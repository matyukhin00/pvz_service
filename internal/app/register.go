package app

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/matyukhin00/pvz_service/internal/model"
)

// @Summary      Регистрация пользователя
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body model.Register true "Данные нового пользователя"
// @Success      201 {object} model.RegisteredUser "Пользователь успешно зарегистрирован"
// @Failure      400 {object} model.Error "Невалидный JSON или ошибка регистрации"
// @Router 		 /register [post]
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
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(model.RegisteredUser{
			Id:    ans.Id,
			Email: ans.Email,
			Role:  ans.Role,
		})
	}
}

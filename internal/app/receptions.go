package app

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/matyukhin00/pvz_service/internal/model"
)

// @Summary      Создание новой приемки товаров (только для сотрудников ПВЗ)
// @Tags         pvz
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body model.Receptions true "ID ПВЗ"
// @Success      201 {object} model.Reception "Приемка создана"
// @Failure      400 {object} model.Error "Невалидный JSON, ПВЗ с заданным id не существует или есть незакрытая приемка"
// @Failure      403 {object} model.Error "Доступ запрещен"
// @Router 		 /receptions [post]
func (s *server) handleReceptions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Context().Value("role") != "employee" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(model.Error{Message: "Access is denied"})
			return
		}

		id := model.Receptions{}
		err := json.NewDecoder(r.Body).Decode(&id)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.Error{Message: "Invalid json format"})
			return
		}

		idString := id.Id.String()
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*3)
		defer cancel()

		exists, err := s.pvzService.Exists(ctx, idString)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.Error{Message: err.Error()})
			return
		}
		if !exists {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.Error{Message: "There is no PVZ with this id"})
			return
		}

		res, err := s.receptionService.Create(ctx, idString)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.Error{Message: err.Error()})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(res)
	}
}

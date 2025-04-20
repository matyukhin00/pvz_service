package app

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/matyukhin00/pvz_service/internal/model"
)

// @Summary      Закрытие последней открытой приемки товаров в рамках ПВЗ (только для сотрудников ПВЗ)
// @Tags         pvz
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param 		 pvzId path string true "UUID ПВЗ" format(uuid)
// @Success      200 {object} model.Reception "Приемка закрыта"
// @Failure      400 {object} model.Error "Невалидный JSON, ПВЗ с заданным id не существует или приемка уже закрыта"
// @Failure      403 {object} model.Error "Доступ запрещен"
// @Router 		 /pvz/{pvzId}/close_last_reception [post]
func (s *server) handleCloseLastReception() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Context().Value("role") != "employee" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(model.Error{Message: "Access is denied"})
			return
		}

		vars := mux.Vars(r)
		pvzIdStr := vars["pvzId"]

		_, err := uuid.Parse(pvzIdStr)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.Error{Message: "Invalid pvzId format"})
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*3)
		defer cancel()
		exists, err := s.pvzService.Exists(ctx, pvzIdStr)
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

		req, err := s.receptionService.Close(ctx, pvzIdStr)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.Error{Message: err.Error()})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(req)

	}
}

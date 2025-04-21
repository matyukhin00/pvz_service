package app

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/matyukhin00/pvz_service/internal/model"
	"github.com/sirupsen/logrus"
)

// @Summary      Создание ПВЗ (только для модераторов)
// @Tags         pvz
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body model.Pvz true "Данные нового ПВЗ"
// @Success      201 {object} model.Pvz "ПВЗ успешно создан"
// @Failure      400 {object} model.Error "Невалидный JSON или ПВЗ уже создан"
// @Failure      403 {object} model.Error "Доступ запрещен"
// @Router		 /pvz [post]
func (s *server) handlePvz() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Context().Value("role") != "moderator" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(model.Error{Message: "Access is denied"})
			return
		}

		pvz := model.Pvz{}

		err := json.NewDecoder(r.Body).Decode(&pvz)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.Error{Message: "Invalid json format"})
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*3)
		defer cancel()

		res, err := s.pvzService.Create(ctx, pvz)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(model.Error{Message: err.Error()})
			return
		}

		go func(id uuid.UUID) {
			s.logger.WithFields(logrus.Fields{
				"method": r.Method,
				"path":   r.URL.Path,
				"user":   r.Context().Value("user_id"),
			}).Infof("PVZ was created with id='%s'", id)
		}(res.Id)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(res)
	}
}

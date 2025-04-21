package app

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/matyukhin00/pvz_service/internal/model"
	"github.com/sirupsen/logrus"
)

// @Summary      Удаление последнего добавленного товара из текущей приемки (LIFO, только для сотрудников ПВЗ)
// @Tags         pvz
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param 		 pvzId path string true "UUID ПВЗ" format(uuid)
// @Success      200 "Товар удален"
// @Failure      400 {object} model.Error "Невалидный JSON, ПВЗ с заданным id не существует или нечего удалять"
// @Failure      403 {object} model.Error "Доступ запрещен"
// @Router 		 /pvz/{pvzId}/delete_last_product [post]
func (s *server) handleDeleteLastProduct() http.HandlerFunc {
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
			json.NewEncoder(w).Encode(model.Error{Message: "There is no PVZ with that id"})
			return
		}

		receptionId, err := s.receptionService.GetId(ctx, pvzIdStr)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.Error{Message: err.Error()})
			return
		}

		err = s.productService.DeleteLast(ctx, receptionId)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.Error{Message: err.Error()})
			return
		}

		res, err := s.receptionService.Get(ctx, receptionId)

		go func(rId, pId uuid.UUID) {
			s.logger.WithFields(logrus.Fields{
				"method": r.Method,
				"path":   r.URL.Path,
				"user":   r.Context().Value("user_id"),
			}).Infof("Last product was deleted from reception (%s) in PVZ='%s'", rId, pId)
		}(res.Id, res.PvzId)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}

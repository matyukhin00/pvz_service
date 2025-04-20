package app

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/matyukhin00/pvz_service/internal/model"
)

// @Summary      Добавление товара в текущую приемку (только для сотрудников ПВЗ)
// @Tags         pvz
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body model.Products true "ID ПВЗ"
// @Success      201 {object} model.Product "Товар добавлен"
// @Failure      400 {object} model.Error "Невалидный JSON, ПВЗ с заданным id не существует или нету открытой приемки"
// @Failure      403 {object} model.Error "Доступ запрещен"
// @Router 		 /products [post]
func (s *server) handleProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Context().Value("role") != "employee" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(model.Error{Message: "Access is denied"})
			return
		}

		request := &model.Products{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.Error{Message: "Invalid json format"})
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*3)
		defer cancel()

		exists, err := s.pvzService.Exists(ctx, request.PvzId.String())
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

		receptionId, err := s.receptionService.GetId(ctx, request.PvzId.String())
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.Error{Message: err.Error()})
			return
		}

		res, err := s.productService.Add(ctx, model.AddProduct{
			Type:        request.Type,
			ReceptionId: receptionId,
		})
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

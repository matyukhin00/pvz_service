package app

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/matyukhin00/pvz_service/internal/model"
)

// @Summary      Получение списка ПВЗ с фильтрацией по дате приемки и пагинацией (для модераторов и сотрудников)
// @Tags         pvz
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param 		 startDate query string true "Начальная дата диапозона (RFC3339)"
// @Param 		 endDate query string true "Конечная дата диапозона (RFC3339)"
// @Param 		 page query integer false "Номер страницы" default(1)
// @Param 		 limit query integer false "Количество элементов на странице" default(10)
// @Success      200 {object} model.PvzInfo "Список ПВЗ"
// @Failure      400 {object} model.Error "Неправильный формат даты или страница пуста"
// @Failure      403 {object} model.Error "Доступ запрещен"
// @Router 		 /pvz [get]
func (s *server) handlePvzPagination() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role := r.Context().Value("role")
		if role != "employee" && role != "moderator" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(model.Error{Message: "Access is denied"})
			return
		}

		query := r.URL.Query()
		startDateStr := query.Get("startDate")
		endDateStr := query.Get("endDate")
		pageStr := query.Get("page")
		limitStr := query.Get("limit")

		startDate, err := time.Parse(time.RFC3339, startDateStr)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.Error{Message: "Invalid format of startDate"})
			return
		}

		endDate, err := time.Parse(time.RFC3339, endDateStr)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.Error{Message: "Invalid format of endDate"})
			return
		}

		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1
		}

		limit, err := strconv.Atoi(limitStr)
		if err != nil || limit < 1 {
			limit = 10
		}

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
		defer cancel()

		pvzListStr, err := s.receptionService.GetFilteredPvz(ctx, startDate, endDate)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.Error{Message: "No info"})
			return
		}

		info := make([]model.PvzInfo, 0)

		for _, pvzStr := range pvzListStr {
			pvz, err := s.pvzService.Get(ctx, pvzStr)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(model.Error{Message: err.Error()})
				return
			}

			pvzInfo := &model.PvzInfo{}
			pvzInfo.Pvz = *pvz
			pvzInfo.Receptions = make([]model.ReceptionInfo, 0)

			receptionList, err := s.receptionService.GetAll(ctx, pvzStr)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(model.Error{Message: err.Error()})
				return
			}

			for _, reception := range receptionList {
				receptionInfo := &model.ReceptionInfo{}
				receptionInfo.Reception = reception

				products, err := s.productService.GetAll(ctx, reception.Id.String())
				if err != nil {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(model.Error{Message: err.Error()})
					return
				}

				receptionInfo.Products = products

				pvzInfo.Receptions = append(pvzInfo.Receptions, *receptionInfo)
			}

			info = append(info, *pvzInfo)
		}

		if page > (len(info)+limit-1)/limit {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.Error{Message: "Page is empty"})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if page*limit == len(info) {
			json.NewEncoder(w).Encode(info[(page-1)*limit:])
		} else {
			json.NewEncoder(w).Encode(info[(page-1)*limit : page*limit])
		}
	}
}

package handler

import (
	"app/app/internal"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type TickerHandler struct {
	service internal.ServiceTicket
}

func NewTickerHandler(service internal.ServiceTicket) *TickerHandler {
	return &TickerHandler{service: service}
}

func (h *TickerHandler) GetTotalAmountTicketsByDestination(w http.ResponseWriter, r *http.Request) {

	dest := chi.URLParam(r, "dest")

	total, err := h.service.GetTicketsAmountByDestinationCountry(dest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{ total: ` + fmt.Sprint(total) + ` }`))

}

func (h *TickerHandler) GetPercentageTicketsByDestination(w http.ResponseWriter, r *http.Request) {

	dest := chi.URLParam(r, "dest")

	percentage, err := h.service.GetPercentageTicketsByDestinationCountry(dest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{ "percentage": "` + fmt.Sprint(percentage) + `%" }`))

}

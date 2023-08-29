package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/WalterPaes/Client-Server-API/internal/exchange/usecases"
)

type Handler struct {
	getExchangeUseCase *usecases.GetExchange
}

func NewHandler(getExchangeUseCase *usecases.GetExchange) *Handler {
	return &Handler{
		getExchangeUseCase: getExchangeUseCase,
	}
}

func (h Handler) GetExchange(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	select {
	case <-time.After(5 * time.Second):
		w.Header().Set("Content-Type", "application/json")
		status := http.StatusOK

		exchange, err := h.getExchangeUseCase.Get(ctx)
		if err != nil {
			status = http.StatusInternalServerError
			w.Write([]byte(err.Error()))
			return
		}
		response, _ := json.Marshal(exchange)

		w.WriteHeader(status)
		w.Write(response)
		return
	case <-ctx.Done():
		log.Println("Request Finalizada pelo Servidor")
	}
}

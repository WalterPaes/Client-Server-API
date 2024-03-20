package handlers

import (
	"context"
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
	ctx := context.Background()

	select {
	case <-time.After(time.Second * 1):
		w.Header().Set("Content-Type", "application/json")

		exchange, err := h.getExchangeUseCase.Get(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		response, _ := json.Marshal(exchange)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
		return

	case <-ctx.Done():
		log.Println("Request Finalizada pelo Servidor")
	}
}

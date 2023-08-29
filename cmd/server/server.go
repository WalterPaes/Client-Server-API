package main

import (
	"github.com/WalterPaes/Client-Server-API/internal/exchange/handlers"
	"github.com/WalterPaes/Client-Server-API/internal/exchange/repository"
	"github.com/WalterPaes/Client-Server-API/internal/exchange/services"
	"github.com/WalterPaes/Client-Server-API/internal/exchange/usecases"
	"github.com/WalterPaes/Client-Server-API/pkg/database"
	"log"
	"net/http"
)

func main() {
	db := database.NewDatabaseConnection()
	defer db.GetConnection().Close()
	db.Migrate()

	getExchangeUC := usecases.NewGetExchange(
		repository.NewRepository(db.GetConnection()),
		services.NewQuotationApi(),
	)

	handler := handlers.NewHandler(getExchangeUC)

	http.HandleFunc("/cotacao", handler.GetExchange)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("SERVER IS DOWN", err)
	}

	log.Println("SERVER IS UP")
}

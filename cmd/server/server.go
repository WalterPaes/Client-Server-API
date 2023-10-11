package main

import (
	"github.com/WalterPaes/Client-Server-API/internal/exchange"
	"github.com/WalterPaes/Client-Server-API/internal/exchange/handlers"
	"github.com/WalterPaes/Client-Server-API/internal/exchange/repository"
	"github.com/WalterPaes/Client-Server-API/internal/exchange/services"
	"github.com/WalterPaes/Client-Server-API/internal/exchange/usecases"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
)

func main() {
	db, err := gorm.Open(sqlite.Open("quotations.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&exchange.Quotation{})
	if err != nil {
		log.Fatal(err)
	}

	getExchangeUC := usecases.NewGetExchange(
		repository.NewRepository(db),
		services.NewQuotationApi(),
	)

	handler := handlers.NewHandler(getExchangeUC)

	http.HandleFunc("/cotacao", handler.GetExchange)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("SERVER IS DOWN", err)
	}

	log.Println("SERVER IS UP")
}

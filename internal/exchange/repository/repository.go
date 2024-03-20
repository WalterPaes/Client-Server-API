package repository

import (
	"context"
	"log"
	"time"

	"github.com/WalterPaes/Client-Server-API/internal/exchange"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) Save(exchange exchange.Quotation) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancel()

	log.Println("Salvando no banco de dados")

	r.db.WithContext(ctx).Create(exchange)

	log.Println("Dados salvos no banco de dados")
}

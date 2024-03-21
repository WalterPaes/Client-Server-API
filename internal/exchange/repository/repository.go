package repository

import (
	"context"
	"fmt"
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

func (r Repository) Save(exchange exchange.Quotation) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancel()

	log.Println("Salvando no banco de dados")

	tx := r.db.WithContext(ctx).Create(exchange)
	if tx.Error != nil {
		return fmt.Errorf("[Repository Error] %s", tx.Error.Error())
	}

	log.Println("Dados salvos no banco de dados")
	return nil
}

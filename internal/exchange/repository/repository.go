package repository

import (
	"context"
	"github.com/WalterPaes/Client-Server-API/internal/exchange"
	"gorm.io/gorm"
	"time"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) Save(parentCtx context.Context, exchange exchange.Quotation) {
	ctx, cancel := context.WithTimeout(parentCtx, time.Millisecond*10)
	defer cancel()

	r.db.WithContext(ctx).Create(exchange)
}

package repository

import (
	"context"
	"database/sql"
	"github.com/WalterPaes/Client-Server-API/internal/exchange"
	"github.com/WalterPaes/Client-Server-API/pkg/customerr"
	"log"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) Save(parentCtx context.Context, exchange exchange.Quotation) error {
	_, cancel := context.WithTimeout(parentCtx, time.Millisecond*10)
	defer cancel()

	query := "INSERT INTO exchanges VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Printf("[QuotationRepository] Message: %s", err.Error())
		return customerr.NewCustomError(err)
	}

	_, err = stmt.Exec(
		exchange.UUID,
		exchange.Code,
		exchange.CodeIn,
		exchange.Name,
		exchange.High,
		exchange.Low,
		exchange.VarBid,
		exchange.PctChange,
		exchange.Bid,
		exchange.Ask,
		exchange.Timestamp,
		exchange.CreateDate,
		exchange.CreatedAt,
	)
	if err != nil {
		log.Printf("[QuotationRepository] Message: %s", err.Error())
		return customerr.NewCustomError(err)
	}

	return nil
}

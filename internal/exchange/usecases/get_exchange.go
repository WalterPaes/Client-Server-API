package usecases

import (
	"context"
	"github.com/WalterPaes/Client-Server-API/internal/exchange"
	"github.com/WalterPaes/Client-Server-API/internal/exchange/adapters"
	"github.com/WalterPaes/Client-Server-API/pkg/customerr"
)

type GetExchange struct {
	quotationRepository     exchange.QuotationRepository
	quotationApiIntegration exchange.QuotationApiIntegration
}

func NewGetExchange(repository exchange.QuotationRepository, apiIntegration exchange.QuotationApiIntegration) *GetExchange {
	return &GetExchange{
		quotationRepository:     repository,
		quotationApiIntegration: apiIntegration,
	}
}

func (g GetExchange) Get(ctx context.Context) (exchange.Exchange, error) {
	var e exchange.Exchange

	quotationResponse, err := g.quotationApiIntegration.Get(ctx)
	if err != nil {
		return e, customerr.NewCustomError(err)
	}

	quotation := adapters.ParseApiResponseToExchange(quotationResponse)

	err = g.quotationRepository.Save(ctx, quotation)
	if err != nil {
		return e, customerr.NewCustomError(err)
	}

	e.CurrentValue = quotation.Bid

	return e, nil
}

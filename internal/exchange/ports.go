package exchange

import (
	"context"
	"github.com/WalterPaes/Client-Server-API/internal/exchange/services"
)

type QuotationRepository interface {
	Save(context.Context, Quotation)
}

type QuotationApiIntegration interface {
	Get(context.Context) (services.QuotationApiResponse, error)
}

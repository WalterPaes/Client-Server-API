package exchange

import (
	"github.com/WalterPaes/Client-Server-API/internal/exchange/services"
)

type QuotationRepository interface {
	Save(Quotation)
}

type QuotationApiIntegration interface {
	Get() (services.QuotationApiResponse, error)
}

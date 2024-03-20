package adapters

import (
	"log"
	"strconv"
	"time"

	"github.com/WalterPaes/Client-Server-API/internal/exchange"
	"github.com/WalterPaes/Client-Server-API/internal/exchange/services"
	"github.com/google/uuid"
)

func ParseApiResponseToExchange(input services.QuotationApiResponse) exchange.Quotation {
	log.Println("Realizando parse do response body para o model")

	high, _ := strconv.ParseFloat(input.Usdbrl.High, 64)
	low, _ := strconv.ParseFloat(input.Usdbrl.Low, 64)
	varBid, _ := strconv.ParseFloat(input.Usdbrl.VarBid, 64)
	pctChange, _ := strconv.ParseFloat(input.Usdbrl.PctChange, 64)
	bid, _ := strconv.ParseFloat(input.Usdbrl.Bid, 64)
	ask, _ := strconv.ParseFloat(input.Usdbrl.Ask, 64)
	createDate, _ := time.Parse(time.RFC3339, input.Usdbrl.CreateDate)

	return exchange.Quotation{
		UUID:       uuid.New().String(),
		Code:       input.Usdbrl.Code,
		CodeIn:     input.Usdbrl.CodeIn,
		Name:       input.Usdbrl.Name,
		High:       high,
		Low:        low,
		VarBid:     varBid,
		PctChange:  pctChange,
		Bid:        bid,
		Ask:        ask,
		Timestamp:  input.Usdbrl.Timestamp,
		CreateDate: createDate,
		CreatedAt:  time.Now(),
	}
}

package services

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/WalterPaes/Client-Server-API/pkg/customerr"
)

type QuotationApi struct{}

func NewQuotationApi() *QuotationApi {
	return &QuotationApi{}
}

func (e QuotationApi) Get(parentCtx context.Context) (QuotationApiResponse, error) {
	ctx, cancel := context.WithTimeout(parentCtx, time.Millisecond*200)
	defer cancel()

	var q QuotationApiResponse
	apiUrl := "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiUrl, nil)
	if err != nil {
		log.Printf("[QuotationApi] Message: %s", err.Error())
		return q, customerr.NewCustomError(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("[QuotationApi] Message: %s", err.Error())
		return q, customerr.NewCustomError(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("[QuotationApi] Message: %s", err.Error())
		return q, customerr.NewCustomError(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var cerr customerr.CustomErr
		err = json.Unmarshal(body, &cerr)
		if err != nil {
			log.Printf("[QuotationApi] Message: %s", err.Error())
			return q, customerr.NewCustomError(err)
		}
		log.Printf("[QuotationApi] Message: %s", cerr.Error())
		return q, cerr
	}

	if err = json.Unmarshal(body, &q); err != nil {
		log.Printf("[QuotationApi] Message: %s", err.Error())
		return q, customerr.NewCustomError(err)
	}

	return q, nil
}

type QuotationApiResponse struct {
	Usdbrl struct {
		Code       string `json:"code"`
		CodeIn     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

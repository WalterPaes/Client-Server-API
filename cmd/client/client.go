package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/WalterPaes/Client-Server-API/internal/exchange"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*300)
	defer cancel()

	apiUrl := "http://localhost:8080/cotacao"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Iniciando Requisição na API de Cotação")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Lendo response body da requisição")
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal("Unexpected Status Code: ", res.StatusCode)
	}

	var e exchange.Exchange
	err = json.Unmarshal(body, &e)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Criando o arquivo de cotação")
	file, err := os.Create("cotacao.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Dólar: %.2f", e.CurrentValue))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Dados salvos no arquivo cotacao.txt")

	select {
	case <-ctx.Done():
		log.Println("Request Finalizada pelo Cliente")
	}
}

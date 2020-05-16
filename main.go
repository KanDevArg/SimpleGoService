package main

import (
	"encoding/json"
	"net/http"

	mw "github.com/kandevarg/SimpleGoService/middleware"
	quotes "github.com/kandevarg/SimpleGoService/quotes"
	log "github.com/op/go-logging"
)

var (
	logger = log.MustGetLogger("quoteservice")
)

func main() {
	quoteService, err := quotes.GetQuoteService()
	if err != nil {
		logger.Fatal(err)
	}

	root := createHTTPHandler(quoteService)
	handler := mw.CombineHandlers(root, mw.RateLimiter(2), mw.LoggingMiddleware)

	logger.Info("Starting server on port 8080...")
	logger.Fatal(http.ListenAndServe(":8080", handler))
}

func createHTTPHandler(quoteService quotes.QuoteService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")

		quote, err := json.Marshal(quoteService.GetRandomQuote())
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			logger.Errorf("marshal error : %v", err)
			return
		}

		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(quote))
	}
}

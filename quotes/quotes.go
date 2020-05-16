package quotes

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
)

// QuoteService interface definition.
type QuoteService interface {
	GetRandomQuote() quote
}

// Service interface implementation for "quotesDataSource" type.
func (qdts quotesDataSource) GetRandomQuote() quote {
	idx := rand.Intn(len(qs.quotes))
	return qdts.quotes[idx]
}

// GetQuoteService provides type that implements Service interface.
func GetQuoteService() (QuoteService, error) {
	b, err := ioutil.ReadFile("data/quotes.json")
	if err != nil {
		return nil, err
	}

	var quotes []quote
	if err := json.Unmarshal(b, &quotes); err != nil {
		return nil, err
	}

	return quotesDataSource{quotes}, nil
}

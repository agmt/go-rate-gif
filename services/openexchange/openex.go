package openexchange

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/shopspring/decimal"
)

type OE struct {
	apiKey string
	base   string
}

type JsonResponse struct {
	Timestamp int
	Base      string
	Rates     map[string]decimal.Decimal
}

func New(apiKey string, base string) *OE {
	return &OE{
		apiKey: apiKey,
		base:   base,
	}
}

func (o *OE) AtDate(symbol string, date time.Time) (decimal.Decimal, error) {
	dateFmt := date.UTC().Format("2006-01-02")
	u, err := url.Parse(fmt.Sprintf("https://openexchangerates.org/api/historical/%s.json", dateFmt))
	if err != nil {
		return decimal.Zero, err
	}

	query := u.Query()
	query.Set("app_id", o.apiKey)
	query.Set("base", o.base)
	query.Add("symbols", symbol)
	u.RawQuery = query.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return decimal.Zero, err
	}

	var j JsonResponse
	err = json.NewDecoder(resp.Body).Decode(&j)
	if err != nil {
		return decimal.Zero, err
	}

	rate, ok := j.Rates[symbol]
	if !ok {
		return decimal.Zero, err
	}

	return rate, nil
}

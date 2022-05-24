package vendorBitcoinRepository

import (
	"encoding/json"
	"fmt"
	"github.com/rcrespodev/Blogs/design/repository/cmd/pkg/domain"
	"net/http"
	"time"
)

type VendorBitcoinRepository struct {
	endpoint string
}

func NewVendorRepository(endpoint string) *VendorBitcoinRepository {
	return &VendorBitcoinRepository{
		endpoint: endpoint,
	}
}

func (v VendorBitcoinRepository) BitcoinPrice() (error, *domain.BitcoinPrice) {
	resp, err := http.Get(v.endpoint)
	if err != nil {
		return fmt.Errorf("internal server error"), nil
	}

	defer resp.Body.Close()

	var VendorBitcoinPrice VendorBitcoinPrice
	err = json.NewDecoder(resp.Body).Decode(&VendorBitcoinPrice)
	if err != nil {
		if err != nil {
			return fmt.Errorf("internal server error"), nil
		}
	}

	return nil, nil
}

func (p VendorBitcoinPrice) newBitcoinPrice(price VendorBitcoinPrice) (*domain.BitcoinPrice, error) {

	rateParser, err := newRateParser(map[string]string{
		price.Bpi.USD.Code: price.Bpi.USD.RateFloat,
		price.Bpi.EUR.Code: price.Bpi.EUR.RateFloat,
		price.Bpi.GBP.Code: price.Bpi.GBP.RateFloat,
	})
	if err != nil {
		return nil, err
	}

	bitcoinPrice := domain.NewBitcoinPrice(time.Now(), []domain.Currency{
		{
			Code:        price.Bpi.USD.Code,
			Rate:        rateParser.GetFloatRate(price.Bpi.USD.Code),
			Description: price.Bpi.USD.Description,
		},
		{
			Code:        price.Bpi.EUR.Code,
			Rate:        rateParser.GetFloatRate(price.Bpi.EUR.Code),
			Description: price.Bpi.EUR.Description,
		},
		{
			Code:        price.Bpi.GBP.Code,
			Rate:        rateParser.GetFloatRate(price.Bpi.GBP.Code),
			Description: price.Bpi.GBP.Description,
		},
	})

	return bitcoinPrice, nil
}

func (v VendorBitcoinRepository) ImplementationName() (error, string) {
	return nil, fmt.Sprintf("Vendor implementation( %v )", v.endpoint)
}
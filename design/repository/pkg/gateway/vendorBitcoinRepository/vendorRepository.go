package vendorBitcoinRepository

import (
	"encoding/json"
	"fmt"
	"github.com/rcrespodev/Blogs/design/repository/pkg/domain"
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

	bitcoinPrice, err := v.newBitcoinPrice(VendorBitcoinPrice)
	if err != nil {
		return err, nil
	}

	return nil, bitcoinPrice
}

func (v VendorBitcoinRepository) newBitcoinPrice(price VendorBitcoinPrice) (*domain.BitcoinPrice, error) {

	bitcoinPrice := domain.NewBitcoinPrice(time.Now(), []domain.Currency{
		{
			Code:        price.Bpi.USD.Code,
			Rate:        price.Bpi.USD.RateFloat,
			Description: price.Bpi.USD.Description,
		},
		{
			Code:        price.Bpi.EUR.Code,
			Rate:        price.Bpi.EUR.RateFloat,
			Description: price.Bpi.EUR.Description,
		},
		{
			Code:        price.Bpi.GBP.Code,
			Rate:        price.Bpi.GBP.RateFloat,
			Description: price.Bpi.GBP.Description,
		},
	})

	return bitcoinPrice, nil
}

func (v VendorBitcoinRepository) ImplementationName() (error, string) {
	return nil, fmt.Sprintf("Vendor implementation( %v )", v.endpoint)
}

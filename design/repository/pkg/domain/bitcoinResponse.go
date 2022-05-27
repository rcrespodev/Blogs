package domain

import "time"

type BitcoinResponse struct {
	BitcoinPrice       *BitcoinPriceResponse `json:"bitcoin_price"`
	ImplementationName string                `json:"implementation_name"`
	Error              string                `json:"error"`
}
type BitcoinPriceResponse struct {
	UpdatedAt  time.Time  `json:"updated_at"`
	CryptoName string     `json:"crypto_name"`
	Currencies []Currency `json:"currencies"`
}

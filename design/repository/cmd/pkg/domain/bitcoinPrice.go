package domain

import "time"

type BitcoinPrice struct {
	updatedAt  time.Time
	cryptoName string
	currencies []Currency
}

type Currency struct {
	Code        string  `json:"code"`
	Rate        float64 `json:"rate"`
	Description string  `json:"description"`
}

func NewBitcoinPrice(updatedAt time.Time, currencies []Currency) *BitcoinPrice {
	//add validation and business rules here
	return &BitcoinPrice{
		updatedAt:  updatedAt,
		cryptoName: "Bitcoin",
		currencies: currencies,
	}
}

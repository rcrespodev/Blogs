package vendorBitcoinRepository

import (
	"strconv"
)

type bitcoinRatesParser struct {
	rates []bitcoinRate
}

type bitcoinRate struct {
	code       string
	stringRate string
	floatRate  float64
}

func newRateParser(vendorRates map[string]string) (*bitcoinRatesParser, error) {
	var b bitcoinRatesParser

	for key, value := range vendorRates {
		b.rates = append(b.rates, bitcoinRate{
			code:       key,
			stringRate: value,
		})
	}

	if err := b.parseBitcoinRate(); err != nil {
		return nil, err
	}

	return &b, nil
}

func (b *bitcoinRatesParser) parseBitcoinRate() error {
	for _, rate := range b.rates {
		if floatRate, err := strconv.ParseFloat(rate.stringRate, 8); err != nil {
			return err
		} else {
			rate.floatRate = floatRate
		}
	}
	return nil
}

func (b bitcoinRatesParser) GetFloatRate(code string) float64 {
	for _, rate := range b.rates {
		if rate.code == code {
			return rate.floatRate
		}
	}
	return 0
}

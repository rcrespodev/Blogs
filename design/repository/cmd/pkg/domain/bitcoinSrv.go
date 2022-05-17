package domain

type BitcoinSrv struct {
	bitcoinRepository BitcoinRepository
}

func NewBitcoinSrv(bitcoinRepository BitcoinRepository) *BitcoinSrv {
	return &BitcoinSrv{bitcoinRepository: bitcoinRepository}
}

func (b BitcoinSrv) GetBitcoinPrice() *BitcoinResponse {
	err, bitcoinPrice := b.bitcoinRepository.BitcoinPrice()
	err, implementationName := b.bitcoinRepository.ImplementationName()

	if err != nil || bitcoinPrice == nil {
		return &BitcoinResponse{
			BitcoinPrice:       nil,
			ImplementationName: "",
			Error:              "internal server error",
		}
	}

	return &BitcoinResponse{
		BitcoinPrice: &BitcoinPriceResponse{
			UpdatedAt:  bitcoinPrice.updatedAt,
			CryptoName: bitcoinPrice.cryptoName,
			Currencies: bitcoinPrice.currencies,
		},
		ImplementationName: implementationName,
		Error:              err.Error(),
	}
}

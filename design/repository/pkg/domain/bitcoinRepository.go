package domain

type BitcoinRepository interface {
	BitcoinPrice() (error, *BitcoinPrice)
	ImplementationName() (error, string)
}

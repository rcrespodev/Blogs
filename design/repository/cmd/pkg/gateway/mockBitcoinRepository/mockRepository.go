package mockBitcoinRepository

import (
	"github.com/rcrespodev/Blogs/design/repository/cmd/pkg/domain"
	"time"
)

type MockBitcoinRepository struct {
	data           *domain.BitcoinPrice
	implementation string
}

func New() domain.BitcoinRepository {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

	return MockBitcoinRepository{
		data: domain.NewBitcoinPrice(t, []domain.Currency{
			{
				Code:        "USD",
				Rate:        29055.3222,
				Description: "United States Dollar",
			},
		}),
		implementation: "Mock_Repository",
	}
}

func (m MockBitcoinRepository) BitcoinPrice() (error, *domain.BitcoinPrice) {
	return nil, m.data
}

func (m MockBitcoinRepository) ImplementationName() (error, string) {
	return nil, m.implementation
}

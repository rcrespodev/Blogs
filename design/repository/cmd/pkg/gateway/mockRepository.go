package gateway

import (
	domain2 "github.com/rcrespodev/Blogs/design/repository/cmd/pkg/domain"
	"time"
)

type MockRepository struct {
	data           *domain2.BitcoinPrice
	implementation string
}

func NewMockRepository() domain2.BitcoinRepository {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

	return MockRepository{
		data: domain2.NewBitcoinPrice(t, []domain2.Currency{
			{
				Code:        "USD",
				Rate:        29055.3222,
				Description: "United States Dollar",
			},
		}),
		implementation: "Mock_Repository",
	}
}

func (m MockRepository) BitcoinPrice() (error, *domain2.BitcoinPrice) {
	return nil, m.data
}

func (m MockRepository) ImplementationName() (error, string) {
	return nil, m.implementation
}

package gateway

import (
	"github.com/rcrespodev/Blogs/design/repository/cmd/pkg/domain"
	"time"
)

type BitcoinRepositoryFactory struct {
	test                                              bool
	mockRepository, redisRepository, vendorRepository domain.BitcoinRepository
}

func NewBitcoinRepositoryFactory(test bool) (error, *BitcoinRepositoryFactory) {
	repo := NewMockRepository()
	return nil, &BitcoinRepositoryFactory{
		test:             test,
		mockRepository:   repo,
		redisRepository:  NewRedisRepository(),
		vendorRepository: NewVendorRepository(),
	}
}

func (b BitcoinRepositoryFactory) Repository() domain.BitcoinRepository {
	if b.test {
		return b.mockRepository
	}

	t := time.Now()
	switch t.Unix() % 2 {
	case 0:
		return b.redisRepository
	default:
		return b.vendorRepository
	}
}

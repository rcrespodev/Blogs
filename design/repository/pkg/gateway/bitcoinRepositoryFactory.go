package gateway

import (
	"github.com/rcrespodev/Blogs/design/repository/pkg/domain"
	"github.com/rcrespodev/Blogs/design/repository/pkg/gateway/mockBitcoinRepository"
	"github.com/rcrespodev/Blogs/design/repository/pkg/gateway/redisBitcoinRepository"
	"github.com/rcrespodev/Blogs/design/repository/pkg/gateway/vendorBitcoinRepository"
	"os"
	"strconv"
	"time"
)

type BitcoinRepositoryFactory struct {
	test                                              bool
	mockRepository, redisRepository, vendorRepository domain.BitcoinRepository
}

func NewBitcoinRepositoryFactory(test bool) (error, *BitcoinRepositoryFactory) {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		return err, nil
	}

	redisRepository, err := redisBitcoinRepository.New(redisHost, redisPort, 0)
	if err != nil {
		return err, nil
	}

	vendorRepository := vendorBitcoinRepository.NewVendorRepository(os.Getenv("VENDOR_ENDPOINT"))

	return nil, &BitcoinRepositoryFactory{
		test:             test,
		mockRepository:   mockBitcoinRepository.New(),
		redisRepository:  redisRepository,
		vendorRepository: vendorRepository,
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

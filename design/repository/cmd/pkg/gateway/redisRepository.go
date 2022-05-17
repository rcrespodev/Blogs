package gateway

import "github.com/rcrespodev/Blogs/design/repository/cmd/pkg/domain"

type RedisRepository struct {
}

func NewRedisRepository() *RedisRepository {
	return &RedisRepository{}
}

func (r RedisRepository) BitcoinPrice() (error, *domain.BitcoinPrice) {
	//TODO implement me
	return nil, nil
}

func (r RedisRepository) ImplementationName() (error, string) {
	//TODO implement me
	return nil, "Redis"
}

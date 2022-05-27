package redisBitcoinRepository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/rcrespodev/Blogs/design/repository/pkg/domain"
	"time"
)

type RedisBitcoinRepository struct {
	redisCliente *redis.Client
	ctx          context.Context
	time         time.Time
}

type RedisBitcoinPrice struct {
	UpdatedAt  time.Time
	CryptoName string
	Currencies []domain.Currency
}

func New(host string, port int, db int) (*RedisBitcoinRepository, error) {
	repository := &RedisBitcoinRepository{
		redisCliente: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%v:%v", host, port),
			Password: "", // no password set
			DB:       db, // use default DB
		}),
		ctx:  context.Background(),
		time: time.Date(2022, time.May, 19, 22, 57, 0, 0, time.UTC),
	}

	BitcoinPrice := domain.NewBitcoinPrice(repository.time, []domain.Currency{
		{
			Code:        "USD",
			Rate:        30185.7069,
			Description: "United States Dollar",
		},
		{
			Code:        "GBP",
			Rate:        24668.4841,
			Description: "British Pound Sterling",
		},
	})

	redisBitcoinPrice := &RedisBitcoinPrice{
		UpdatedAt:  BitcoinPrice.UpdatedAt(),
		CryptoName: BitcoinPrice.CryptoName(),
		Currencies: BitcoinPrice.Currencies(),
	}

	marshalBitcoinPrice, err := json.Marshal(redisBitcoinPrice)
	if err != nil {
		return nil, err
	}

	if _, err := repository.redisCliente.Do(repository.ctx, "set", repository.time.String(), marshalBitcoinPrice).Result(); err != nil {
		return nil, err
	}

	return repository, nil
}

func (r *RedisBitcoinRepository) BitcoinPrice() (error, *domain.BitcoinPrice) {
	result, err := r.redisCliente.Do(r.ctx, "get", r.time.String()).Result()
	if err != nil {
		return err, nil
	}

	strResult, ok := result.(string)
	if !ok {
		return fmt.Errorf("internal Server Error"), nil
	}

	var redisBitcoinPrice RedisBitcoinPrice
	err = json.Unmarshal([]byte(strResult), &redisBitcoinPrice)
	//bitcoinPrice, ok := result.(*domain.BitcoinPrice)
	if err != nil {
		return fmt.Errorf("internal Server Error"), nil
	}

	return nil, domain.NewBitcoinPrice(redisBitcoinPrice.UpdatedAt, redisBitcoinPrice.Currencies)
}

func (r RedisBitcoinRepository) ImplementationName() (error, string) {
	return nil, "Redis"
}

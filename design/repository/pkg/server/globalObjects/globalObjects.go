package globalObjects

import (
	"fmt"
	"github.com/rcrespodev/Blogs/design/repository/pkg/gateway"
	"os"
)

var Factory *gateway.BitcoinRepositoryFactory

func New() error {
	err, f := gateway.NewBitcoinRepositoryFactory(false)
	if err == nil {
		Factory = f
	}

	envs := []string{
		"APP_HOST", "APP_PORT", "REDIS_HOST", "REDIS_PORT", "VENDOR_ENDPOINT",
	}

	for _, env := range envs {
		if envValue := os.Getenv(env); envValue == "" {
			err = fmt.Errorf("env %v not found", env)
			break
		}
	}

	return err
}

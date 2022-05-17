package globalObjects

import (
	"github.com/rcrespodev/Blogs/design/repository/cmd/pkg/gateway"
)

var Factory *gateway.BitcoinRepositoryFactory

func New() error {
	err, f := gateway.NewBitcoinRepositoryFactory(false)
	if err == nil {
		Factory = f
	}

	return err
}

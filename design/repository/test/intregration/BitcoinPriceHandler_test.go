package intregration

import (
	"github.com/rcrespodev/Blogs/design/repository/api/v1/handlers"
	"github.com/rcrespodev/Blogs/design/repository/pkg/domain"
	"github.com/rcrespodev/Blogs/design/repository/pkg/server/globalObjects"
	"log"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestBitcoinSrvHandler(t *testing.T) {
	tests := []struct {
		name string
		resp *domain.BitcoinResponse
	}{
		{
			name: "Base Test",
			resp: &domain.BitcoinResponse{
				BitcoinPrice:       nil,
				ImplementationName: "",
				Error:              "",
			},
		},
		// TODO: Add test cases.
	}

	err := os.Setenv("REDIS_HOST", "localhost")
	err = os.Setenv("REDIS_PORT", "6379")
	err = os.Setenv("VENDOR_ENDPOINT", "https://api.coindesk.com/v1/bpi/currentprice.json")
	err = os.Setenv("APP_HOST", "localhost")
	err = os.Setenv("APP_PORT", "8080")
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := globalObjects.New(); err != nil {
		log.Fatal(err.Error())
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 10; i++ {
				_, response := handlers.HttpGetBitcoinPrice()
				resp, ok := response.(*domain.BitcoinResponse)
				if !ok {
					log.Fatal(ok)
				}

				if !reflect.DeepEqual(resp.Error, tt.resp.Error) {
					t.Errorf("\n- Implementation: %v\n - actual Error:\n\t %v\n- expected Error:\n\t %v", resp.ImplementationName, resp.Error, tt.resp.Error)
				}
				time.Sleep(time.Microsecond * 100000)
			}
		})
	}
}

package intregration

import (
	"github.com/rcrespodev/Blogs/design/repository/cmd/pkg/domain"
	"github.com/rcrespodev/Blogs/design/repository/cmd/pkg/server/globalObjects"
	"github.com/rcrespodev/Blogs/design/repository/cmd/pkg/server/handlers"
	"log"
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
					t.Errorf("\n- actual Error:\n\t %v\n- expected Error:\n\t %v", resp.Error, tt.resp.Error)
				}
				time.Sleep(time.Microsecond * 100000)
			}
		})
	}
}

package unit

import (
	"github.com/rcrespodev/Blogs/design/repository/cmd/pkg/domain"
	"github.com/rcrespodev/Blogs/design/repository/cmd/pkg/gateway/mockBitcoinRepository"
	"reflect"
	"testing"
	"time"
)

func TestBitcoinSrv_GetBitcoinPrice(t *testing.T) {
	type args struct {
		repository domain.BitcoinRepository
	}

	type resp struct {
		resp *domain.BitcoinResponse
	}

	tests := []struct {
		name string
		args args
		resp resp
	}{
		{
			name: "base Test",
			args: args{repository: mockBitcoinRepository.New()},
			resp: resp{resp: &domain.BitcoinResponse{
				BitcoinPrice: &domain.BitcoinPriceResponse{
					UpdatedAt:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					CryptoName: "Bitcoin",
					Currencies: []domain.Currency{
						{
							Code:        "USD",
							Rate:        29055.3222,
							Description: "United States Dollar",
						},
					},
				},
				ImplementationName: "Mock_Repository",
			},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := domain.NewBitcoinSrv(tt.args.repository)
			if resp := srv.GetBitcoinPrice(); !reflect.DeepEqual(resp, tt.resp.resp) {
				t.Errorf("\n- actual Response:\n\t %v\n- expected Response:\n\t %v", resp, tt.resp.resp)
			}
		})
	}
}

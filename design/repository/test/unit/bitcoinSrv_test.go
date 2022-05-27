package unit

import (
	domain2 "github.com/rcrespodev/Blogs/design/repository/pkg/domain"
	"github.com/rcrespodev/Blogs/design/repository/pkg/gateway/mockBitcoinRepository"
	"reflect"
	"testing"
	"time"
)

func TestBitcoinSrv_GetBitcoinPrice(t *testing.T) {
	type args struct {
		repository domain2.BitcoinRepository
	}

	type resp struct {
		resp *domain2.BitcoinResponse
	}

	tests := []struct {
		name string
		args args
		resp resp
	}{
		{
			name: "base Test",
			args: args{repository: mockBitcoinRepository.New()},
			resp: resp{resp: &domain2.BitcoinResponse{
				BitcoinPrice: &domain2.BitcoinPriceResponse{
					UpdatedAt:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					CryptoName: "Bitcoin",
					Currencies: []domain2.Currency{
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
			srv := domain2.NewBitcoinSrv(tt.args.repository)
			if resp := srv.GetBitcoinPrice(); !reflect.DeepEqual(resp, tt.resp.resp) {
				t.Errorf("\n- actual Response:\n\t %v\n- expected Response:\n\t %v", resp, tt.resp.resp)
			}
		})
	}
}

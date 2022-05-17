package gateway

import "github.com/rcrespodev/Blogs/design/repository/cmd/pkg/domain"

type VendorRepository struct {
}

func NewVendorRepository() *VendorRepository {
	return &VendorRepository{}
}

func (v VendorRepository) BitcoinPrice() (error, *domain.BitcoinPrice) {
	//TODO implement me
	return nil, nil
}

func (v VendorRepository) ImplementationName() (error, string) {
	//TODO implement me
	return nil, "Vendor"
}

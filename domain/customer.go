package domain

import "context"

type Customer struct {
	CustomerID          int64
	CompanyName         string
	FisrtName           string
	LastName            string
	BillingAddress      string
	City                string
	StateOrProvince     string
	ZIPCode             string
	Email               string
	CompanyWebsite      string
	PhoneNumber         string
	FaxNumber           string
	ShipAddress         string
	ShipCity            string
	ShipStateOrProvince string
	ShipZIPCode         string
	ShipPhoneNumber     string
}

type CustomerRepository interface {
	Save(ctx context.Context, c []*Customer) error
}

type CustomerUsecase interface {
	Save(ctx context.Context, c []*Customer) error
}

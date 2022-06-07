package models

type Address struct {
	Id      int64
	address string
}

type AddressRetrieve struct {
	Address *Address
	Error   error
}

package models

type Address struct {
	Id      int64
	Address string
}

type AddressRetrieve struct {
	Address *Address
	Error   error
}

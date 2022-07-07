package models

const (
	Available string = "available"
	Moved            = "moved"
	Issued           = "issued"
)

type OrderAvailability struct {
	OrderId      int64
	IssuePointId int64
	Status       string
}

type OrderAvailabilityRetrieve struct {
	OrderAvailability *OrderAvailability
	Error             error
}

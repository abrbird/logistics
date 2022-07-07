package service

type Service interface {
	Address() AddressService
	IssuePoint() IssuePointService
	OrderAvailability() OrderAvailabilityService
}

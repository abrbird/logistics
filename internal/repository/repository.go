package repository

type Repository interface {
	Address() AddressRepository
	IssuePoint() IssuePointRepository
	OrderAvailability() OrderAvailabilityRepository
}

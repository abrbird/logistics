package implemented_service

import "gitlab.ozon.dev/zBlur/homework-3/logistics/internal/service"

type Service struct {
	addressService           *AddressService
	issuePointService        *IssuePointService
	orderAvailabilityService *OrderAvailabilityService
}

func New() *Service {
	return &Service{
		addressService:           &AddressService{},
		issuePointService:        &IssuePointService{},
		orderAvailabilityService: &OrderAvailabilityService{},
	}
}

func (s *Service) Address() service.AddressService {
	return s.addressService
}

func (s *Service) IssuePoint() service.IssuePointService {
	return s.issuePointService
}

func (s *Service) OrderAvailability() service.OrderAvailabilityService {
	return s.orderAvailabilityService
}

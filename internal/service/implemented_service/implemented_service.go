package implemented_service

import (
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/cache"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/service"
)

type Service struct {
	cache                    cache.Cache
	addressService           *AddressService
	issuePointService        *IssuePointService
	orderAvailabilityService *OrderAvailabilityService
}

func New(cache cache.Cache) *Service {
	srvc := &Service{
		cache: cache,
	}
	srvc.addressService = &AddressService{srvc}
	srvc.issuePointService = &IssuePointService{srvc}
	srvc.orderAvailabilityService = &OrderAvailabilityService{srvc}

	return srvc
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

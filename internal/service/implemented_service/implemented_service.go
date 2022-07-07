package implemented_service

import (
	"github.com/abrbird/logistics/internal/cache"
	"github.com/abrbird/logistics/internal/service"
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

package sql_repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/repository"
)

type SQLRepository struct {
	dbConnectionPool *pgxpool.Pool

	addressRepository           *SQLAddressRepository
	issuePointRepository        *SQLIssuePointRepository
	orderAvailabilityRepository *SQLOrderAvailabilityRepository
}

func New(dbConnPool *pgxpool.Pool) *SQLRepository {
	repo := &SQLRepository{
		dbConnectionPool: dbConnPool,
	}
	repo.addressRepository = &SQLAddressRepository{store: repo}
	repo.issuePointRepository = &SQLIssuePointRepository{store: repo}
	repo.orderAvailabilityRepository = &SQLOrderAvailabilityRepository{store: repo}

	return repo
}

func (s *SQLRepository) Address() repository.AddressRepository {
	return s.addressRepository
}

func (s *SQLRepository) IssuePoint() repository.IssuePointRepository {
	return s.issuePointRepository
}

func (s *SQLRepository) OrderAvailability() repository.OrderAvailabilityRepository {
	return s.orderAvailabilityRepository
}

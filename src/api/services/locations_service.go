package services

import (
	"fmt"
	"testing/src/api/domain/locations"
	locationsprovider "testing/src/api/providers/locations_provider"
	"testing/src/api/utils/errors"
)

type locationsService struct{}
type locationsServiceInterface interface {
	GetCountry(countryId string) (*locations.Country, *errors.ApiError)
}

var (
	LocationsService locationsServiceInterface
)

func init() {
	fmt.Println("init service")
	LocationsService = &locationsService{}
}

func (s *locationsService) GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	fmt.Println("inside service")
	return locationsprovider.GetCountry(countryId)
}

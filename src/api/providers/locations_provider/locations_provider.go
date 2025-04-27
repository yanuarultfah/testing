package locationsprovider

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing/src/api/domain/locations"
	"testing/src/api/utils/errors"

	"github.com/federicoleon/golang-restclient/rest"
	// "github.com/mercadolibre/golang-restclient/rest"
)

const (
	urlGetCountry = "https://api.mercadolibre.com/countries/%s"
)

func GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	fmt.Println(fmt.Sprintf(urlGetCountry, countryId))
	response := rest.Get(fmt.Sprintf(urlGetCountry, countryId))
	fmt.Println(response.String())
	fmt.Println(fmt.Sprintf("obtained status code:%d", response.StatusCode))
	if response == nil || response.Response == nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("invalid restclient response when trying to get country %s", countryId),
		}
	}

	if response.StatusCode > 299 {
		var apiErr errors.ApiError
		if err := json.Unmarshal(response.Bytes(), &apiErr); err != nil {
			return nil, &errors.ApiError{
				Status:  http.StatusInternalServerError,
				Message: fmt.Sprintf("invalid error interface when getting country %s", countryId),
			}
		}
		return nil, &apiErr
	}

	var result locations.Country
	if err := json.Unmarshal(response.Bytes(), &result); err != nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("error when trying unmarshall data for %s", countryId),
		}
	}

	return &result, nil
}

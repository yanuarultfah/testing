package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"testing/src/api/utils/errors"

	"github.com/federicoleon/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestGetCountriesNotFound(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:        "https://api.mercadolibre.com/countries/AR",
		HTTPMethod: http.MethodGet,
		// RespHTTPCode: http.StatusOK,
		// RespBody:     `{"id": "AR", "name": "Argentina", "locale": "es-AR", "currency_id": "ARS", "time_zone": "GMT-3", "area_code": 54, "region_id": "LAC", "region_name": "Latin America and the Caribbean", "country_code": 1, "country_name": "Argentina"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "country not found", "error": "not_found", "status": 404}`,
	})
	response, err := http.Get("http://localhost:8080/locations/countries/AR")
	assert.Nil(t, err)
	assert.NotNil(t, response)
	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))

	var apiErr errors.ApiError
	err = json.Unmarshal(bytes, &apiErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
	assert.EqualValues(t, "not_found", apiErr.Error)
	assert.EqualValues(t, "country not found", apiErr.Message)
}

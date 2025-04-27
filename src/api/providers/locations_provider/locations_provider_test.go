package locationsprovider

import (
	"net/http"
	"os"
	"testing"

	"github.com/federicoleon/golang-restclient/rest"

	// "github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

// var mock = flag.Bool("mock", false, "Use mock server")

// func TestMain(m *testing.M) {
// 	flag.Parse()
// 	if *mock {
// 		rest.StartMockupServer()
// 	}
// 	os.Exit(m.Run())
// }

func TestGetCountryRestclientError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
	})

	country, err := GetCountry("AR")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient error when getting country AR", err.Message)
}

func TestGetCountryNotFound(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "country not found", "error": "not_found", "status": 404}`,
	})

	country, err := GetCountry("AR")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "country not found", err.Message)
}

func TestGetCountryInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusInternalServerError,
		RespBody:     `{"message": 123, "error": "server_error", "status": 500}`,
	})

	country, err := GetCountry("AR")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when getting country AR", err.Message)
}

func TestGetCountryInvalidJsonResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": "123", "name": "Argentina", "time_zone": "GMT-03:00"}`,
	})

	country, err := GetCountry("AR")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshall getting country AR", err.Message)
}

func TestGetCountryNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": "AR", "name": "ARGENTINA", "time_zone": "GMT-3", "states": [{"id": "BA", "name": "Buenos Aires"}]}`,
	})

	country, err := GetCountry("AR")
	assert.Nil(t, err)
	assert.NotNil(t, country)
	assert.EqualValues(t, "AR", country.Id)
	assert.EqualValues(t, "ARGENTINA", country.Name)
	assert.EqualValues(t, "GMT-3", country.TimeZone)
	assert.EqualValues(t, 1, len(country.States))
}

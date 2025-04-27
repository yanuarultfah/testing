package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"testing/src/api/domain/locations"
	"testing/src/api/services"
	"testing/src/api/utils/errors"

	"github.com/federicoleon/golang-restclient/rest"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	getCountryFunc func(countryId string) (*locations.Country, *errors.ApiError)
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

type locationsServiceMock struct {
}

func (*locationsServiceMock) GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	return getCountryFunc(countryId)
}

func TestGetCountryNoError(t *testing.T) {
	getCountryFunc = func(countryId string) (*locations.Country, *errors.ApiError) {
		return &locations.Country{Id: "AR", Name: "Argentinasaf"}, nil
	}
	services.LocationsService = &locationsServiceMock{}
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "country_id", Value: "AR"},
	}
	GetCountry(c)
	assert.EqualValues(t, http.StatusOK, response.Code)
	var country locations.Country
	// var apiErr *errors.ApiError
	err := json.Unmarshal(response.Body.Bytes(), &country)
	// fmt.Println(string(response.Body.Bytes()))
	assert.Nil(t, err)
	assert.EqualValues(t, "AR", country.Id)
	assert.EqualValues(t, "Argentina", country.Name)
}

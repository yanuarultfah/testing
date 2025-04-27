package app

import "testing/src/api/controller"

func mapUrls() {
	router.GET("/locations/countries/:country_id", controller.GetCountry)
}

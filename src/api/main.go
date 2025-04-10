package main

import (
	"fmt"
	locationsprovider "testing/src/api/providers/locations_provider"
)

func main() {
	country, err := locationsprovider.GetCountry("AR")
	fmt.Println(err)
	fmt.Println(country)
}

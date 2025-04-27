package test

import (
	"fmt"
	"os"
	"testing"
	"testing/src/api/app"

	"github.com/federicoleon/golang-restclient/rest"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	fmt.Println("about to start the application")
	go app.StartApp()
	fmt.Println("application started,about to start tests cases")
	os.Exit(m.Run())
}

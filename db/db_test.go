package db

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestConnectNoConnection(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("r should not be nil")
		}
	}()

	vars := []byte("MONGO_URI=test")
	err := ioutil.WriteFile(".env", vars, 0644)

	if err != nil {
		panic(err)
	}

	Connect()
}

func TestConnectNoEnvVariable(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("r should not be nil")
		}
	}()

	err := os.Remove(".env")

	if err != nil {
		panic(err)
	}

	Connect()
}

package main

import (
	"os"
	"testing"
)

//TestParseConfigFilter verify that json file can be read
func TestParseConfig(t *testing.T) {
	var err error

	_, err = ParseConfig()
	if err != nil {
		t.Fatal(err)
	}

}

//TestDefaultchecker verify logic and functionality of valuechecker
func TestDefaultchecker(t *testing.T) {
	var err error

	_, err = Defaultchecker("test","test","test")
	if err != nil {
		t.Fatal(err)
	}

	_, err = Defaultchecker("","test","test")
	if err != nil {
		t.Fatal(err)
	}

	_, err = Defaultchecker("test","","test")
	if err != nil {
		t.Fatal(err)
	}
	
	_, err = Defaultchecker("","","test")
	if err == nil {
		t.Fatal(err)
	}

	_, err = Defaultchecker("test","test","")
	if err == nil {
		t.Fatal(err)
	}

}

//TestMain run presteps before m.Run and teardown after
func TestMain(m *testing.M) {
	output := m.Run()
	os.Exit(output)
}

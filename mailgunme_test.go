package main

import (
	"log"
	"os"
	"testing"
)

//TestParseConfigFilter verify that json file can be read
func TestParseConfig(t *testing.T) {
	ckey := os.Getenv("USER")
	os.Setenv("USER", "null")
	ParseConfig()

}

//TestMain run presteps before m.Run and teardown after
func TestMain(m *testing.M) {
	output := m.Run()
	os.Exit(output)
}

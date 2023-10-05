package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup()

	os.Exit(m.Run())
}

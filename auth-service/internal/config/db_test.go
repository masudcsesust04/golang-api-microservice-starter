package config

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	cleanup := TestDBCleaner(m)
	defer cleanup()
	os.Exit(m.Run())
}

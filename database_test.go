package main

import (
	"testing"
)

func TestWriteKeyValue(t *testing.T) {
	loadConfigGlobals()
	pool := newDbPool()

	writeKeyValue(pool, "testkey", "testvalue")
}

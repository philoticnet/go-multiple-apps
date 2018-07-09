package main

import (
	"testing"
)

func TestMain(t *testing.T) {
}

func TestMain2(t *testing.T) {
	t.Fatalf("App2 main failed")
}

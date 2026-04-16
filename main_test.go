package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	got := add(1, 2)
	want := 3
	assert.Equal(t, got, want)
}

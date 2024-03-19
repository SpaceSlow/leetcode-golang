package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsHappy19(t *testing.T) {
	num := 19
	expected := true

	assert.Equal(t, isHappy(num), expected)
}

func TestIsHappy2(t *testing.T) {
	num := 2
	expected := false

	assert.Equal(t, isHappy(num), expected)
}

func TestIsHappy4(t *testing.T) {
	num := 4
	expected := false

	assert.Equal(t, isHappy(num), expected)
}

func TestIsHappy1(t *testing.T) {
	num := 1
	expected := true

	assert.Equal(t, isHappy(num), expected)
}

package add_digits

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddDigits38(t *testing.T) {
	num := 38
	expected := 2

	assert.Equal(t, addDigits(num), expected)
}
func TestAddDigits123(t *testing.T) {
	num := 123
	expected := 6

	assert.Equal(t, addDigits(num), expected)
}
func TestAddDigits9(t *testing.T) {
	num := 9
	expected := 9

	assert.Equal(t, addDigits(num), expected)
}
func TestAddDigits27(t *testing.T) {
	num := 27
	expected := 9

	assert.Equal(t, addDigits(num), expected)
}
func TestAddDigits0(t *testing.T) {
	num := 0
	expected := 0

	assert.Equal(t, addDigits(num), expected)
}

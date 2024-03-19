package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTimeRequiredToBuy1(t *testing.T) {
	tickets := []int{2, 3, 2}
	k := 2
	expected := 6

	assert.Equal(t, expected, timeRequiredToBuy(tickets, k))
}

func TestTimeRequiredToBuy2(t *testing.T) {
	tickets := []int{5, 1, 1, 1}
	k := 0
	expected := 8

	assert.Equal(t, expected, timeRequiredToBuy(tickets, k))
}

func TestTimeRequiredToBuy3(t *testing.T) {
	tickets := []int{84, 49, 5, 24, 70, 77, 87, 8}
	k := 3
	expected := 154

	assert.Equal(t, expected, timeRequiredToBuy(tickets, k))
}

func TestTimeRequiredToBuy4(t *testing.T) {
	tickets := []int{15, 66, 3, 47, 71, 27, 54, 43, 97, 34, 94, 33, 54, 26, 15, 52, 20, 71, 88, 42, 50, 6, 66, 88, 36, 99, 27, 82, 7, 72}
	k := 18
	expected := 1457

	assert.Equal(t, expected, timeRequiredToBuy(tickets, k))
}

func TestTimeRequiredToBuy5(t *testing.T) {
	tickets := []int{5, 1, 4, 3, 5, 1}
	k := 3
	expected := 13

	assert.Equal(t, expected, timeRequiredToBuy(tickets, k))
}

func TestTimeRequiredToBuy100(t *testing.T) {
	tickets := make([]int, 100)
	for i := range tickets {
		tickets[i] = 100
	}
	k := 99
	expected := 10000

	assert.Equal(t, expected, timeRequiredToBuy(tickets, k))
}

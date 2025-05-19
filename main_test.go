package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// generator tests
func TestGenerateRandomElements(t *testing.T) {

	tests := []struct {
		size     int
		sizeWant int
	}{
		{-5, 0},                  //negative
		{0, 0},                   //empty
		{5, 5},                   //small size
		{10_000_000, 10_000_000}, //large size
	}

	for _, test := range tests {
		data := generateRandomElements(test.size)
		assert.Len(t, data, test.sizeWant)
	}
}

// Maximum
func TestMaximum(t *testing.T) {
	tests := []struct {
		data    []int
		maxWant int
	}{
		{[]int{}, 0},
		{[]int{1000}, 1000},
		{[]int{3, 5, 10, -20, 0, 100, -4}, 100},
	}
	for _, test := range tests {
		assert.Equal(t, test.maxWant, maximum(test.data))
	}

}

// Maximum with many flows
func TestMaxChunks(t *testing.T) {

	tests := []struct {
		data    []int
		maxWant int
	}{
		{[]int{}, 0},
		{[]int{1000}, 1000},
		{[]int{3, 5, 10, -20, 0, 100, -4, 1, 1, 1, 1}, 100},
	}

	for _, test := range tests {
		assert.Equal(t, test.maxWant, maxChunks(test.data))
	}
}

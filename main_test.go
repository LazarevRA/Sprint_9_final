package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// generator tests
func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		size    int
		wantErr bool
	}{
		{0, true},            //zero case
		{-10, true},          //negative case
		{10, false},          //small size
		{100_000_000, false}, //large size
	}

	for _, test := range tests {
		data, err := generateRandomElements(test.size)

		if test.wantErr {
			require.Error(t, err)
			require.Nil(t, data)
		} else {
			require.NoError(t, err)
			require.NotNil(t, data)
			assert.Len(t, data, test.size)
		}

	}
}

// Maximum
func TestMaximumEmptySlice(t *testing.T) {
	_, err := maximum([]int{})
	require.Error(t, err)
}

func TestMaximumSingleElement(t *testing.T) {
	res, err := maximum([]int{1})
	require.NoError(t, err)
	assert.Equal(t, 1, res)
}

func TestMaximumOK(t *testing.T) {
	res, err := maximum([]int{1, 2, 3, 10, 0})
	require.NoError(t, err)
	assert.Equal(t, 10, res)
}

// Maximum with many flows
func TestMaxChunksEmpty(t *testing.T) {
	_, err := maxChunks([]int{})
	require.Error(t, err)
}
func TestMaxChunksSmallData(t *testing.T) {
	data := []int{1, 2, 3, 10}
	max, err := maxChunks(data)
	require.NoError(t, err)
	assert.Equal(t, 10, max)
}

func TestMaxChunksOK(t *testing.T) {
	data := make([]int, 10000)
	for i := range data {
		data[i] = i
	}
	maxWant := 111_111
	data[10] = maxWant
	max, err := maxChunks(data)
	require.NoError(t, err)
	assert.Equal(t, maxWant, max)
}

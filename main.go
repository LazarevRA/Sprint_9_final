package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		fmt.Println("slice size must have positive integer value")
		return nil
	}

	data := make([]int, size)
	for i := range data {
		data[i] = rand.Int()
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		fmt.Println("empty data slice")
		return 0
	}
	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	var wg sync.WaitGroup

	if len(data) == 0 {
		fmt.Println("can't find maximum in empty slice")
		return 0
	}

	if len(data) < CHUNKS {
		fmt.Println("data slice size is smaller than the number of chunks. It is recommended to use search of maximum conducted in one thread")
	}

	sizeChunk := len(data) / CHUNKS
	if sizeChunk == 0 {
		sizeChunk = 1
	}

	maxVals := make([]int, CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * sizeChunk
		end := start + sizeChunk

		if i == CHUNKS-1 {
			end = len(data)
		}

		if start >= len(data) {
			break
		}

		wg.Add(1)
		go func(chunk []int, i int) {
			defer wg.Done()
			locMax := maximum(chunk)

			maxVals[i] = locMax

		}(data[start:end], i)
	}
	wg.Wait()

	absMax := maximum(maxVals)

	return absMax
}

func main() {

	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	startMax := time.Now()

	fmt.Println("Ищем максимальное значение в один поток")

	max := maximum(data)

	elapsedMax := time.Since(startMax).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsedMax)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	startMaxChunks := time.Now()
	maxChunks := maxChunks(data)

	elapsedMaxChunks := time.Since(startMaxChunks).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxChunks, elapsedMaxChunks)
}

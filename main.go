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
func generateRandomElements(size int) ([]int, error) {
	if size <= 0 {
		err := fmt.Errorf("slice size must have positive integer value")
		return nil, err
	}

	data := make([]int, size)
	for i := range data {
		data[i] = rand.Int()
	}
	return data, nil
}

// maximum returns the maximum number of elements.
func maximum(data []int) (int, error) {
	if len(data) == 0 {
		return 0, fmt.Errorf("can't find maximum in empty slice")
	}
	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max, nil
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) (int, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	if len(data) == 0 {
		return 0, fmt.Errorf("can't find maximum in empty slice")
	}

	if len(data) < CHUNKS {
		fmt.Println("data slice size is smaller than the number of chunks. It is recommended to use search of maximum conducted in one thread")
	}

	sizeChunk := len(data) / CHUNKS
	if sizeChunk == 0 {
		sizeChunk = 1
	}

	maxVals := make([]int, 0, CHUNKS)

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
		go func(chunk []int) {
			defer wg.Done()
			locMax, err := maximum(chunk)
			if err != nil {
				fmt.Println("error in finding maximum with gorutines", err)
				return
			}
			mu.Lock()
			maxVals = append(maxVals, locMax)
			mu.Unlock()
		}(data[start:end])
	}
	wg.Wait()

	absMax, err := maximum(maxVals)

	if err != nil {
		return 0, fmt.Errorf("error in finding maximum in maxSlice: %w", err)
	}

	return absMax, nil
}

func main() {

	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data, err := generateRandomElements(SIZE)
	if err != nil {
		fmt.Println(err)
		return
	}

	startMax := time.Now()

	fmt.Println("Ищем максимальное значение в один поток")

	max, err := maximum(data)

	if err != nil {
		fmt.Println("Ошибка поиска максимума в один поток:", err)
	}

	elapsedMax := time.Since(startMax).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsedMax)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	startMaxChunks := time.Now()
	maxChunks, err := maxChunks(data)
	if err != nil {
		fmt.Println("Ошибка поиска максимума в несколько потоков:", err)
	}

	elapsedMaxChunks := time.Since(startMaxChunks).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxChunks, elapsedMaxChunks)
}

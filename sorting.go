package main

import (
	"sort"
	"time"
)

// Function to sort a single array
func sortArray(arr []int) {
	sort.Ints(arr)
}

// Function to sort all arrays sequentially
func sequentialSorting(arrays [][]int) time.Duration {
	startTime := time.Now()

	// Implement sequential sorting logic

	return time.Since(startTime)
}

// Function to sort a single array concurrently
func sortArrayConcurrently(arr []int, ch chan<- []int) {
	sort.Ints(arr)
	ch <- arr
}

// Function to sort all arrays concurrently
func concurrentSorting(arrays [][]int) time.Duration {
	startTime := time.Now()

	// Implement concurrent sorting logic

	return time.Since(startTime)
}

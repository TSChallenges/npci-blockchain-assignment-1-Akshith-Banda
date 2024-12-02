package main

import (
	"errors"
	"fmt"
	"sort"
)

// FilterAndSort filters numbers greater than or equal to the threshold and sorts them.
func FilterAndSort(nums []int, threshold int) []int {
	// TODO: Implement this function
	sliceToSort := []int{}
	for _, num := range nums {
		if num >= threshold {
			sliceToSort = append(sliceToSort, num)
		}
	}

	sort.Ints(sliceToSort)
	return sliceToSort
}

// FindMostFrequent finds the most frequent word in a slice of strings.
// Returns an error if the slice is empty.
func FindMostFrequent(words []string) (string, error) {
	// TODO: Implement this function

	if len(words) == 0 {
		return "", errors.New("slice is empty")
	}
	wordsMap := make(map[string]int)
	for _, word := range words {
		wordsMap[word]++
	}

	var frequentWord string
	var frequentWordCount int
	for k, v := range wordsMap {
		if v > frequentWordCount {
			frequentWord = k
			frequentWordCount = v
		}
	}
	return frequentWord, nil
}

func main() {
	// Test FilterAndSort
	fmt.Println("Testing FilterAndSort:")
	nums := []int{3, 10, 1, 7, 8, 2}
	threshold := 5
	fmt.Printf("Input: %v, Threshold: %d, Output: %v\n", nums, threshold, FilterAndSort(nums, threshold))

	// Test FindMostFrequent
	fmt.Println("\nTesting FindMostFrequent:")
	words := []string{"apple", "banana", "apple", "orange", "banana", "apple"}
	result, err := FindMostFrequent(words)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Input: %v, Most Frequent: %s\n", words, result)
	}

	// Edge case: empty slice
	emptyWords := []string{}
	result, err = FindMostFrequent(emptyWords)
	if err != nil {
		fmt.Printf("Error: %v (empty input case handled)\n", err)
	} else {
		fmt.Printf("Input: %v, Most Frequent: %s\n", emptyWords, result)
	}
}
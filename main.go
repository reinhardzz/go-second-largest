package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

func findSecondLargest(arr []int) (int, error) {
	n := len(arr)
	if n < 2 {
		return 0, fmt.Errorf("array must contain at least two elements")
	}

	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	return arr[1], nil
}

func secondLargestHandler(w http.ResponseWriter, r *http.Request) {
	values, ok := r.URL.Query()["numbers"]
	if !ok || len(values[0]) < 1 {
		http.Error(w, "URL parameter 'numbers' is missing", http.StatusBadRequest)
		return
	}

	var numbers []int
	for _, numStr := range values {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			http.Error(w, "Invalid number format in 'numbers' parameter", http.StatusBadRequest)
			return
		}
		numbers = append(numbers, num)
	}

	secondLargest, err := findSecondLargest(numbers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := map[string]int{"second_largest": secondLargest}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func findMostDuplicate(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("array must not be empty")
	}

	// Count occurrences of each number in the array
	counts := make(map[int]int)
	for _, num := range numbers {
		counts[num]++
	}

	// Find the number with the maximum count (most duplicated)
	maxCount := 0
	mostDuplicate := 0
	for num, count := range counts {
		if count > maxCount {
			maxCount = count
			mostDuplicate = num
		}
	}

	return mostDuplicate, nil
}

func mostDuplicateHandler(w http.ResponseWriter, r *http.Request) {
	values, ok := r.URL.Query()["numbers"]
	if !ok || len(values[0]) < 1 {
		http.Error(w, "URL parameter 'numbers' is missing", http.StatusBadRequest)
		return
	}

	var numbers []int
	for _, numStr := range values {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			http.Error(w, "Invalid number format in 'numbers' parameter", http.StatusBadRequest)
			return
		}
		numbers = append(numbers, num)
	}

	mostDuplicate, err := findMostDuplicate(numbers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := map[string]int{"most_duplicate": mostDuplicate}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/second-largest", secondLargestHandler).Methods("GET")
	r.HandleFunc("/api/most-duplicate", mostDuplicateHandler).Methods("GET")

	fmt.Println("API server listening on :8080")
	http.ListenAndServe(":8080", r)
}

package main

import "log"

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 100, 232, 654, 657, 1000}

	for index, value := range nums {
		if index != binarySearch(nums, value) {
			log.Fatal("error")
		}
	}
}

func binarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		middle := (low + high) / 2

		if target < nums[middle] {
			high = middle - 1
		} else if target > nums[middle] {
			low = middle + 1
		} else {
			return middle
		}
	}
	return -1
}

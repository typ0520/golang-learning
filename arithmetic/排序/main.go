package main

import (
	"fmt"
	"log"
)

// https://leetcode-cn.com/problems/sort-an-array/solution/tu-jie-jing-dian-pai-xu-by-wang-41/
func main() {
	//定长的数组传递到函数中时是值传递，如果是变长数组传递的是引用类型
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	bubbleSort(nums)
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] < nums[i-1] {
			log.Fatal("error")
		}
	}

	nums2 := []int{4, 3, 8, 5, 2, 6, 1, 7}
	selectSort(nums2)
	for i := 0; i < len(nums2); i++ {
		if i != 0 && nums2[i] < nums2[i-1] {
			log.Fatal("error")
		}
	}

	nums3 := []int{4, 3, 8, 5, 2, 6, 1, 7}
	insertSort(nums3)
	for i := 0; i < len(nums3); i++ {
		if i != 0 && nums3[i] < nums3[i-1] {
			log.Fatal("error")
		}
	}

	nums4 := []int{47, 29, 71, 99, 78, 19, 24, 47}
	quickSort(nums4, 0, len(nums4)-1)
	for i := 0; i < len(nums4); i++ {
		if i != 0 && nums4[i] < nums4[i-1] {
			log.Fatal("error")
		}
	}
}

// 冒泡排序 O(n*n)
func bubbleSort(nums []int) {
	fmt.Println("-----------bubbleSort-----------")
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		fmt.Println(nums)
	}
}

// 选择排序 O(n*n)
func selectSort(nums []int) {
	fmt.Println("-----------selectSort-----------")
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		minValueIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minValueIndex] {
				minValueIndex = j
			}
		}
		if minValueIndex != i {
			nums[i], nums[minValueIndex] = nums[minValueIndex], nums[i]
		}
		fmt.Println(nums)
	}
}

// 插入排序
func insertSort(nums []int) {
	fmt.Println("-----------insertSort-----------")
	fmt.Println(nums)
	for i := 1; i < len(nums); i++ {
		j := i
		tempValue := nums[i]
		for j > 0 && nums[j-1] > tempValue {
			nums[j] = nums[j-1]
			j--
		}
		nums[j] = tempValue
		fmt.Println(nums)
	}
}

// 快速排序
func quickSort(nums []int, l int, r int) {
	fmt.Println("-----------quickSort-----------")
	fmt.Println(nums)
	// base := nums[l]
	// quickSort(nums, l, middle)
	// quickSort(nums, middle, r)
}

// 堆排序
// 归并排序
// 希尔排序

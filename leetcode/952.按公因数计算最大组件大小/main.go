package main

import (
	"fmt"
	"log"
)

func main() {
	// 给定一个由不同正整数的组成的非空数组 A，考虑下面的图：

	// 有 A.length 个节点，按从 A[0] 到 A[A.length - 1] 标记；
	// 只有当 A[i] 和 A[j] 共用一个大于 1 的公因数时，A[i] 和 A[j] 之间才有一条边。
	// 返回图中最大连通组件的大小。

	//

	// 示例 1：

	// 输入：[4,6,15,35]
	// 输出：4

	// 示例 2：

	// 输入：[20,50,9,63]
	// 输出：2

	// 示例 3：

	// 输入：[2,3,6,7,4,12,21,39]
	// 输出：8

	// 来源：力扣（LeetCode）
	// 链接：https://leetcode-cn.com/problems/largest-component-size-by-common-factor
	// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

	// nums1 := []int{4, 6, 15, 35}
	// result1 := largestComponentSize(nums1)
	// if result1 != 4 {
	// 	log.Fatal(fmt.Sprintf("error1: %d", result1))
	// }

	nums2 := []int{20, 50, 9, 63}
	result2 := largestComponentSize(nums2)
	if result2 != 2 {
		log.Fatal(fmt.Sprintf("error2: %d", result2))
	}

	// nums3 := []int{2, 3, 6, 7, 4, 12, 21, 39}
	// result3 := largestComponentSize(nums3)
	// if result3 != 8 {
	// 	log.Fatal(fmt.Sprintf("error3: %d", result3))
	// }
}

func largestComponentSize(A []int) int {
	table := make(map[int]int)
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if getGreatestCommonFactor(A[i], A[j]) > 1 {
				min, max := A[j], A[i]
				if A[i] < A[j] {
					min, max = A[i], A[j]
				}
				count, ok := table[min]
				if ok {
					table[max] = count + 1
					delete(table, min)
				} else {
					table[max] = 2
				}
			}
		}
	}
	fmt.Println(table)

	count := 0
	for _, v := range table {
		count += v
	}
	return count
}

func getGreatestCommonFactor(a int, b int) int {
	fmt.Printf("%d, %d", a, b)

	if b > a {
		a, b = b, a
	}
	for b != 0 {
		multiplier := a / b
		remainder := a % b
		a = (a - remainder) / multiplier
		b = remainder
	}
	fmt.Printf(", %d\n", a)
	return a
}

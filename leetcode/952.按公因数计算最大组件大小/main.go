package main

import (
	"fmt"
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

	table := make(map[int]*[]int)
	fmt.Println(table)

	//head tail len

	// var arr = [...]int{5:2}
	// //数组指针
	// var pf *[6]int = &arr
	// (*pf)[0] = 1
	// fmt.Println(arr)

	// arr := make([]int, 3)

	// table[0] = &arr

	// var arr2 []int
	// arr2, _ = table[0]

	// arr2[0] = 1
	// fmt.Println(arr2)

	// fmt.Printf("p: %p, T: %T\n", &arr, &arr)
	// fmt.Println(len(arr))
	// fmt.Println(arr)
	// fmt.Println(table)

	nums1 := []int{4, 6, 15, 35}
	result1 := largestComponentSize2(nums1)
	fmt.Println(result1)
	// if result1 != 4 {
	// 	log.Fatal(fmt.Sprintf("error1: %d", result1))
	// }

	nums2 := []int{20, 50, 9, 63}
	result2 := largestComponentSize2(nums2)
	fmt.Println(result2)

	// if result2 != 2 {
	// 	log.Fatal(fmt.Sprintf("error2: %d", result2))
	// }

	nums3 := []int{2, 3, 6, 7, 4, 12, 21, 39}
	result3 := largestComponentSize2(nums3)
	fmt.Println(result3)
	// if result3 != 8 {
	// 	log.Fatal(fmt.Sprintf("error3: %d", result3))
	// }
}

func largestComponentSize2(A []int) int {
	type node struct {
		headVal int
		tailVal int
		length  int
	}
	table := make(map[int][]node)

	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if getGreatestCommonFactor(A[i], A[j]) > 1 {
				min, max := A[i], A[j]
				if min > max {
					min, max = max, min
				}
				v, ok := table[min]
				if ok {
					var array []node
					for _, n := range v {
						array = append(array, node{
							headVal: n.headVal,
							tailVal: max,
							length:  n.length + 1,
						})
					}
					table[max] = append(table[max], array...)
					//fmt.Println(array)
				} else {
					var array []node
					array = append(array, node{
						headVal: min,
						tailVal: max,
						length:  2,
					})
					table[max] = array
				}
			}
		}
	}
	length := 0
	for k, array := range table {
		fmt.Printf("%d ", k)
		fmt.Println(array)
		for _, n := range array {
			if n.length > length {
				length = n.length
			}
		}
	}
	return length
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
	count := 0
	for _, v := range table {
		count += v
	}
	return count
}

func getGreatestCommonFactor(a int, b int) int {
	if b > a {
		a, b = b, a
	}
	//c, d := a, b
	for b != 0 {
		multiplier := a / b
		remainder := a % b
		a = (a - remainder) / multiplier
		b = remainder
	}
	// if a > 1 {
	// 	fmt.Printf("%d %d", d, c)
	// 	fmt.Printf(" , %d\n", a)
	// }

	return a
}

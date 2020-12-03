package main

import "fmt"

func main() {
	// 给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。

	// 进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？

	//

	// 示例 1：

	// 输入：nums1 = [1,3], nums2 = [2]
	// 输出：2.00000
	// 解释：合并数组 = [1,2,3] ，中位数 2
	// 示例 2：

	// 输入：nums1 = [1,2], nums2 = [3,4]
	// 输出：2.50000
	// 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
	// 示例 3：

	// 输入：nums1 = [0,0], nums2 = [0,0]
	// 输出：0.00000
	// 示例 4：

	// 输入：nums1 = [], nums2 = [1]
	// 输出：1.00000
	// 示例 5：

	// 输入：nums1 = [2], nums2 = []
	// 输出：2.00000
	//

	// 提示：

	// nums1.length == m
	// nums2.length == n
	// 0 <= m <= 1000
	// 0 <= n <= 1000
	// 1 <= m + n <= 2000
	// -106 <= nums1[i], nums2[i] <= 106

	// 来源：力扣（LeetCode）
	// 链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
	// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

	// array := [...][2]int{
	// 	{1, 2},
	// 	{3, 4},
	// }
	// //fmt.Println(array)
	// excepted := [...]float32{2.5}
	// fmt.Println(excepted)

	// for i, v := range array {
	// 	findMedianSortedArrays(array[i], array[i])
	// 	fmt.Println(v)
	// }

	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var array []int
	if len(nums1) == 0 {
		array = nums2
	} else if len(nums2) == 0 {
		array = nums1
	} else if nums1[len(nums1)-1] <= nums2[0] {
		array = append(nums1, nums2...)
	} else if nums1[0] >= nums2[len(nums2)-1] {
		array = append(nums2, nums1...)
	} else {
		var (
			index  int = len(nums1) + len(nums2) - 1
			index1 int = len(nums1) - 1
			index2 int = len(nums2) - 1
		)
		array = make([]int, index+1)

		for index1 >= 0 || index2 >= 0 {
			if index1 >= 0 && index2 >= 0 {
				if nums1[index1] >= nums2[index2] {
					array[index] = nums1[index1]
					index1--
				} else {
					array[index] = nums2[index2]
					index2--
				}

			} else if index1 >= 0 {
				array[index] = nums1[index1]
				index1--
			} else if index2 >= 0 {
				array[index] = nums2[index2]
				index2--
			}
			index--
		}
	}
	if len(array)%2 == 0 {
		n1 := array[len(array)/2]
		n2 := array[len(array)/2-1]
		return float64(n1+n2) / 2.0
	}
	return float64(array[len(array)/2])
}

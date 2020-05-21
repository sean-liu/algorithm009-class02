package solutions

// https://leetcode-cn.com/problems/merge-sorted-array/

// 三指针法，一个是1号数组最末尾的指针，第二个是1号数组的最后一个元素，第三个是2号数组的最后一个元素，依次往前推
// 时间复杂度 O(m+n), 空间复杂度 O(1)
func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1Index := m - 1
	nums2Index := n - 1
	lastFreeIndex := m + n - 1

	for ; nums1Index >= 0 && nums2Index >= 0; lastFreeIndex-- {
		if nums1[nums1Index] >= nums2[nums2Index] {
			nums1[lastFreeIndex] = nums1[nums1Index]
			if nums1Index >= 0 {
				nums1Index--
			}
		} else {
			nums1[lastFreeIndex] = nums2[nums2Index]
			if nums2Index >= 0 {
				nums2Index--
			}
		}
	}

	// 多余的2号数组元素要copy到1号数组去
	for ; nums2Index >= 0; nums2Index-- {
		nums1[nums2Index] = nums2[nums2Index]
	}
}

// 暴力法，把2号数组copy过去，然后sort，时间 O((m+n)log(m+n)), 空间 O(1)
// 双指针法，开个新的数组，依次从1号2号数组中拿过去。时间O(m+n), 空间 O(m+n)

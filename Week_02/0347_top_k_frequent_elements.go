package solutions

import "container/heap"

// https://leetcode-cn.com/problems/top-k-frequent-elements/

// use heap, count then fetch, time O(n), space O(n)
func topKFrequent(nums []int, k int) []int {
	frequenceMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		_, found := frequenceMap[nums[i]]
		if !found {
			frequenceMap[nums[i]] = 0
		}
		frequenceMap[nums[i]]++
	}

	maxHeap := &TopKHeap{}
	for k, v := range frequenceMap {
		maxHeap.Push(&TopKHeapItem{Key: k, Value: v})
	}
	heap.Init(maxHeap)

	result := make([]int, 0)
	for i := k; i > 0; i-- {
		result = append(result, (heap.Pop(maxHeap).(*TopKHeapItem)).Key)
	}

	return result
}

// TopKHeapItem used in heap
type TopKHeapItem struct {
	Key   int
	Value int
}

// TopKHeap heap for int
type TopKHeap []*TopKHeapItem

func (h TopKHeap) Len() int           { return len(h) }
func (h TopKHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h TopKHeap) Less(i, j int) bool { return h[i].Value > h[j].Value }

// Pop pop function, required by heap
func (h *TopKHeap) Pop() interface{} {
	result := (*h)[h.Len()-1]
	*h = (*h)[0 : h.Len()-1]
	return result
}

// Push push function, required by heap
func (h *TopKHeap) Push(val interface{}) {
	*h = append(*h, val.(*TopKHeapItem))
}

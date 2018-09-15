package sort

import (
	"github.com/rodolfo-r/datastructures"
	"github.com/rodolfo-r/datastructures/container/priorityqueue/binaryheap"
)

// Heap applies heapsort to nums
func Heap(nums []int) {
	var pq datastructures.PriorityQueue = binaryheap.BuildMax(nums)

	for ok := true; ok; _, ok = pq.Dequeue() {
	}
}

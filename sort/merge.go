package sort

import (
	"github.com/rodolfo-r/datastructures/container/queue/doublyqueue"
)

// Merge applies merge sort to nums.
func Merge(nums []int) {
	merge(nums, 0, len(nums)-1)
}

func merge(nums []int, start, end int) {
	if start >= end {
		return
	}
	mid := (end + start) / 2
	merge(nums, start, mid)
	merge(nums, mid+1, end)
	mergeSorted(nums, start, mid, end)
}

func mergeSorted(nums []int, start, mid, end int) {
	q, q2 := doublyqueue.New(), doublyqueue.New()
	for i := start; i <= mid; i++ {
		q.Enqueue(nums[i])
	}

	for i := mid + 1; i <= end; i++ {
		q2.Enqueue(nums[i])
	}

	i := start
	for !(q.Size() == 0) && !(q2.Size() == 0) {
		if q.Front().(int) < q2.Front().(int) {
			n, ok := q.Dequeue()
			if !ok {
				break
			}

			nums[i], i = n.(int), i+1

		} else {
			n, ok := q2.Dequeue()
			if !ok {
				break
			}

			nums[i], i = n.(int), i+1
		}
	}

	for q.Size() > 0 {
		n, ok := q.Dequeue()
		if !ok {
			break
		}

		nums[i], i = n.(int), i+1
	}

	for q2.Size() > 0 {
		n, ok := q2.Dequeue()
		if !ok {
			break
		}

		nums[i], i = n.(int), i+1
	}
}

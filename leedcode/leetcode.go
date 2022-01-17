package leetcode

import (
	"math/rand"
)

// 27. RemoveElement
func removeElement(nums []int, val int) int {
	n := len(nums)
	for i := 0; i < n; {
		if nums[i] == val {
			n--
			nums[i] = nums[n]
		} else {
			i++
		}
	}
	return n
}

// 2. TwoSum
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		idx, ok := m[target-v]
		if ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil
}

// 382.  Linked List Random Node
type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	count int
	head  *ListNode
	pool  []int
}

func Constructor(head *ListNode) Solution {
	var s Solution
	s.head = head
	return s
}

func (this *Solution) GetRandom() int {
	r := rand.Intn(this.count + 1)
	this.count++
	if r == this.count-1 && this.head != nil {
		t := this.head
		this.head = this.head.Next
		this.pool = append(this.pool, t.Val)
		return t.Val
	} else {
		return this.pool[r%len(this.pool)]
	}

}

package solutions

import "testing"

func TestMajorityElement(t *testing.T) {
	result := majorityElement([]int{2, 2, 1, 1, 1, 2, 2})
	assertEqual(t, result, 2)
}

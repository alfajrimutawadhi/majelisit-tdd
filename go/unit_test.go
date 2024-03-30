package unit

import (
	"reflect"
	"testing"
)

func Sum(a, b int) int {
	// write code here
	return 0
}

func Test_Sum(t *testing.T) {
	tests := []struct {
		name string
		num1 int
		num2 int
		want int
	}{
		{
			name: "test 1",
			num1: 1,
			num2: 2,
			want: 3,
		},
		{
			name: "test 2",
			num1: 1,
			num2: 5,
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.num1, tt.num2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Failed; got : %d, want : %d", got, tt.want)
				return
			}
		})
	}
}

// Question:
// Given an array of integers nums and an integer target,
// return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution,
// and you may not use the same element twice.

// Example:
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

func TwoSum(nums []int, target int) []int {
	// write code here
	return nil
}

func Test_TwoSum(t *testing.T) {
	tests := []struct {
		name   string
		in     []int
		target int
		out    []int
	}{
		{
			name:   "test 1",
			in:     []int{2, 7, 11, 15},
			target: 9,
			out:    []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.in, tt.target)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("Failed; got : %d, want : %d", got, tt.out)
				return
			}
		})
	}
}

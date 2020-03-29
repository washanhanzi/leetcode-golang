package twoSum

import "testing"

func TestTwoSum(t *testing.T) {
	got := twoSum([]int{2, 7, 11, 15}, 9)
	want := []int{0, 1}

	verify := make(map[int]bool)
	for _, d := range want {
		verify[d] = true
	}
	for _, d := range got {
		if !verify[d] {
			t.Errorf("got %q want %q", got, want)
		}
	}
}

package wxclient

import (
	"testing"
)

func TestCode2Session(t *testing.T) {
	f := func(p []int) int {
		return 1
	}
	t.Log(f([]int{1, 2, 3, 4}))
}

package utils

import "testing"

func TestMap(t *testing.T) {
	src := []int{1, 2, 3, 4, 5}
	fn := func(i int) int {
		return i + 1
	}
	dst := Map(src, fn)
	t.Log(dst)
}

func TestFilter(t *testing.T) {
	src := []int{1, 2, 3, 4, 5}
	fn := func(i int) bool {
		return i > 3
	}
	dst := Filter(src, fn)
	t.Log(dst)
}

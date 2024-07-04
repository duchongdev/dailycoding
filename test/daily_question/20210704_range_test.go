package test

import "testing"

func TestForRange(t *testing.T) {

	s := []int{1, 2, 3}

	for i := range s {
		s = append(s, i)
	}

	t.Logf("s = %v", s)

}

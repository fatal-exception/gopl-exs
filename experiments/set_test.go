package set

import (
	"fmt"
	"testing"
)

func IntSliceEquals(t *testing.T, s1 []int, s2 []int) bool {
	t.Helper()
	if len(s1) != len(s2) {
		return false
	}
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}

func TestSet(t *testing.T) {
	t.Run("Set can be added to", func(t *testing.T) {
		s := New()
		s.Add(6)
		got := s.Contains(6)
		if got != true {
			t.Errorf("Expected true, got %v", got)
		}
	})

	t.Run("Set contains is accurate", func(t *testing.T) {
		s := New()
		s.Add(6)
		yes := s.Contains(6)
		no := s.Contains(0)
		if yes != true {
			t.Errorf("yes expected true, got %v", yes)
		}
		if no != false {
			t.Errorf("no expected false, got %v", no)
		}
	})

	t.Run("Set can be removed from", func(t *testing.T) {
		s := New()
		s.Add(6)
		got := s.Contains(6)
		if got != true {
			t.Errorf("Expected true, got %v", got)
		}
		s.Remove(6)
		got = s.Contains(6)
		if got != false {
			t.Errorf("Expected false, got %v", got)
		}
	})
	t.Run("Set can give length", func(t *testing.T) {
		s := New()
		s.Add(6)
		s.Add(2)
		s.Add(4)
		len := s.Len()
		if len != 3 {
			t.Errorf("Len: expected 3, got %v", len)
		}
	})
	t.Run("Set equality", func(t *testing.T) {
		s1 := New()
		s1.Add(6)
		s1.Add(2)
		s1.Add(4)
		s2 := New()
		s2.Add(6)
		s2.Add(2)
		s2.Add(4)
		resTrue := s1.Equals(s2)
		s3 := New()
		resFalse := s1.Equals(s3)
		if !resTrue {
			t.Errorf("Expected res to be true, was %v", resTrue)
		}
		if resFalse {
			t.Errorf("Expected res to be false, was %v", resFalse)
		}
	})

	t.Run("Set varargs add", func(t *testing.T) {
		s := New()
		s.Add(6, 4, 2)
		for _, val := range []int{6, 4, 2} {
			res := s.Contains(val)
			if res != true {
				t.Errorf("Expected %d to be present, but got %v", val, res)
			}
		}
	})
	t.Run("Set union", func(t *testing.T) {
		s1 := New()
		s1.Add(1, 2, 3)
		s2 := New()
		s2.Add(2, 4, 6)
		s3 := s1.Union(s2)
		s3Length := s3.Len()
		expectedLength := 5
		if s3Length != expectedLength {
			t.Errorf("Expected %d but got %d from s3", expectedLength, s3Length)
		}
	})
	t.Run("Set to slice", func(t *testing.T) {
		s1 := New()
		s1.Add(2, 1, 3, 2)
		slice := s1.ToSlice()
		expected := []int{1, 2, 3}
		if !IntSliceEquals(t, slice, expected) {
			t.Errorf("map slice %v did not match expected %v", slice, expected)
		}
	})
	t.Run("Set intersection", func(t *testing.T) {
		s1 := New()
		s1.Add(1, 2, 3)
		s2 := New()
		s2.Add(2, 4, 6)
		s3 := s1.Intersect(s2)
		s3Length := s3.Len()
		expectedLength := 1
		if s3Length != expectedLength {
			t.Errorf("Expected %d but got %d from s3", expectedLength, s3Length)
		}
	})
	t.Run("Constructor takes args", func(t *testing.T) {
		s1 := New(1, 2, 3)
		expectedSet := New()
		expectedSet.Add(1)
		expectedSet.Add(2)
		expectedSet.Add(3)
		if !s1.Equals(expectedSet) {
			t.Errorf("Expected %v to match %v", s1, expectedSet)
		}
	})
	t.Run("Decent string output", func(t *testing.T) {
		s1 := New(4, 5, 8)
		stringOutput := fmt.Sprint(s1)
		expectedOutput := "Set[4, 5, 8]"
		if stringOutput != expectedOutput {
			t.Errorf("Expected %s but got %s", expectedOutput, stringOutput)
		}
	})
	t.Run("Set subtraction", func(t *testing.T) {
		s1 := New(1, 2, 3, 4)
		s2 := New(2, 3)
		s3 := s1.Subtract(s2)
		expected := New(1, 4)
		if !s3.Equals(expected) {
			t.Errorf("Set s3 after subtraction expected %v, got %v", expected, s3)
		}
	})
}

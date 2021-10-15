package skiplist

import "testing"

func TestSkipList(t *testing.T) {
	sl := NewSkipList()

	sl.Insert("a", 1)
	sl.String()
}

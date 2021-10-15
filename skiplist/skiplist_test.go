package skiplist

import "testing"

func TestSkipList(t *testing.T) {
	sl := NewSkipList()

	sl.Insert("a", 1)
	sl.String()
	t.Log(sl)
	t.Log(sl.head.forwards[0])

	sl.Insert("b", 2)
	sl.String()
	t.Log(sl)
	t.Log(sl.head.forwards[0])
	t.Log(sl.head.forwards[0].forwards[0])

}

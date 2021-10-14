package skiplist

import "testing"

func TestSkiplist(t *testing.T) {
	sl := NewSkipList(1)
	sl.Add(12, "twelve")
	sl.Add(11, "eleven")
	t.Logf("now skip list: %d", sl.Len())
	t.Logf("get %d from skip : %+v", 11, sl.Search(11))
	sl.Remove(11)
	t.Logf("now skip list: %d", sl.Len())
	t.Logf("get %d from skip : %+v", 11, sl.Search(11))
}

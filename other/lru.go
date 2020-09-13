package main

/*
最近最久未使用（LRU）

他的实质是，当需要置换一页时，选择最近一段时间里最久没有使用过的页面予以置换。这种算法就称为最久未使用算法（Last Recently Used,LRU）。
*/

func lru(key string) {

	var lMap = make(map[string]int)
	// 计数
	if item, ok := lMap[key]; ok {
		lMap[key] = item + 1
	} else {
		lMap[key] = 1
	}

}

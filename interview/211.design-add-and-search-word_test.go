package interview

import "testing"

func TestWordDictionary211(t *testing.T) {
	wd := NewWordDictionary()

	wd.AddWord("bad")
	wd.AddWord("dad")
	wd.AddWord("mad")

	// Search 带 '.' 通配符：搜索 ".ad" 就是找所有3字母词且以ad结尾
	if !wd.Search(".ad") {
		t.Error(".ad should match bad, dad, mad")
	}
	// 搜索 "b.." 找所有3字母词且以b开头
	if !wd.Search("b..") {
		t.Error("b.. should match bad")
	}
	// 精确匹配
	if !wd.Search("bad") {
		t.Error("bad should match")
	}
	// 不存在的精确词
	if wd.Search("pad") {
		t.Error("pad should not exist")
	}
}

func TestWordDictionaryEmpty(t *testing.T) {
	wd := NewWordDictionary()
	if wd.Search("a") {
		t.Error("empty dict should not contain 'a'")
	}
}

func TestWordDictionaryDotOnly(t *testing.T) {
	wd := NewWordDictionary()
	wd.AddWord("a")
	if !wd.Search(".") {
		t.Error(". should match a")
	}
}

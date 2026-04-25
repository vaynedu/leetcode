package interview

import "testing"

func TestLRUCacheInterface(t *testing.T) {
	// 用接口测试，避免构造函数名冲突
	cache := Constructor146(2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	if cache.Get(1) != 1 {
		t.Errorf("Get(1) should be 1, got %d", cache.Get(1))
	}

	cache.Put(3, 3) // 淘汰 key=2

	if cache.Get(2) != -1 {
		t.Errorf("Get(2) should be -1, got %d", cache.Get(2))
	}

	cache.Put(4, 4) // 淘汰 key=1

	if cache.Get(1) != -1 {
		t.Errorf("Get(1) should be -1, got %d", cache.Get(1))
	}
	if cache.Get(3) != 3 {
		t.Errorf("Get(3) should be 3, got %d", cache.Get(3))
	}
	if cache.Get(4) != 4 {
		t.Errorf("Get(4) should be 4, got %d", cache.Get(4))
	}
}

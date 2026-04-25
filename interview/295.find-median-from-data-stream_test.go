package interview

import "testing"

func TestMedianFinder(t *testing.T) {
	m := Constructor295()
	
	// 按 LeetCode 示例顺序
	m.AddNum(1)
	m.AddNum(2)
	if m.FindMedian() != 1.5 {
		t.Errorf("after 1,2 median = %v, want 1.5", m.FindMedian())
	}
	
	m.AddNum(3)
	if m.FindMedian() != 2.0 {
		t.Errorf("after 1,2,3 median = %v, want 2.0", m.FindMedian())
	}
	
	// 奇数个测试
	m2 := Constructor295()
	m2.AddNum(5)
	if m2.FindMedian() != 5.0 {
		t.Errorf("single element median = %v, want 5.0", m2.FindMedian())
	}
	
	// 偶数个测试
	m3 := Constructor295()
	m3.AddNum(1)
	m3.AddNum(2)
	if m3.FindMedian() != 1.5 {
		t.Errorf("1,2 median = %v, want 1.5", m3.FindMedian())
	}
}

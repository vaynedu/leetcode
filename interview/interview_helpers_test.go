package interview

// 测试辅助函数（同一 package 下所有 _test.go 共享）

func intEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func intSliceEq(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !intEq(a[i], b[i]) {
			return false
		}
	}
	return true
}

func strEq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[string]bool)
	for _, v := range a {
		m[v] = true
	}
	for _, v := range b {
		if !m[v] {
			return false
		}
	}
	return true
}

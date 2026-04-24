package interview

// ================================================================
// 56. 合并区间
// ================================================================

// Interval 区间
type Interval struct {
	Start int
	End   int
}

// merge 合并区间
func merge(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return nil
	}

	// 1. 按起始位置排序
	sortByStart(intervals)

	// 2. 遍历合并
	result := []Interval{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := &result[len(result)-1]
		curr := intervals[i]

		if curr.Start <= last.End {
			// 重叠，合并
			if curr.End > last.End {
				last.End = curr.End
			}
		} else {
			// 不重叠，直接加入
			result = append(result, curr)
		}
	}
	return result
}

// sortByStart 按起始位置排序
func sortByStart(intervals []Interval) {
	for i := 0; i < len(intervals)-1; i++ {
		for j := i + 1; j < len(intervals); j++ {
			if intervals[i].Start > intervals[j].Start {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}
}

// intervalsEqual 比较两个区间列表是否相等
func intervalsEqual(a, b []Interval) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Start != b[i].Start || a[i].End != b[i].End {
			return false
		}
	}
	return true
}

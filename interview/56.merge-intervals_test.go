package interview

import "testing"

func TestMergeIntervalsSingleProblemCases(t *testing.T) {
	tests := []struct {
		name      string
		intervals []Interval
		want      []Interval
	}{
		{
			name:      "基本重叠",
			intervals: []Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want:      []Interval{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "端点相接也合并",
			intervals: []Interval{{1, 4}, {4, 5}},
			want:      []Interval{{1, 5}},
		},
		{
			name:      "乱序输入先排序",
			intervals: []Interval{{8, 10}, {1, 3}, {15, 18}, {2, 6}},
			want:      []Interval{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "完全包含",
			intervals: []Interval{{1, 10}, {2, 3}, {4, 8}},
			want:      []Interval{{1, 10}},
		},
		{
			name:      "负数区间",
			intervals: []Interval{{-10, -1}, {-3, 2}, {5, 7}},
			want:      []Interval{{-10, 2}, {5, 7}},
		},
		{
			name:      "单个区间",
			intervals: []Interval{{1, 4}},
			want:      []Interval{{1, 4}},
		},
		{
			name:      "空输入",
			intervals: []Interval{},
			want:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := merge(tt.intervals)
			if !intervalsEqual(got, tt.want) {
				t.Fatalf("merge() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

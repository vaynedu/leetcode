package interview

import "testing"

func TestKSmallestPairs(t *testing.T) {
	tests := []struct {
		name   string
		nums1  []int
		nums2  []int
		k      int
		want   int // 最少对数
	}{
		{
			name:   "基本",
			nums1:  []int{1, 7, 11},
			nums2:  []int{2, 4, 6},
			k:      3,
			want:   3,
		},
		{
			name:   "k小于可用对数",
			nums1:  []int{1, 1, 2},
			nums2:  []int{1, 2, 3},
			k:      2,
			want:   2,
		},
		{
			name:   "k为0",
			nums1:  []int{1, 2},
			nums2:  []int{3},
			k:      0,
			want:   0,
		},
		{
			name:   "空数组",
			nums1:  []int{},
			nums2:  []int{1},
			k:      3,
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kSmallestPairs(tt.nums1, tt.nums2, tt.k)
			if len(got) != tt.want {
				t.Errorf("kSmallestPairs(...) returned %d pairs, want %d", len(got), tt.want)
			}
			// 验证每对的 sum 是升序的
			for i := 1; i < len(got); i++ {
				sumPrev := got[i-1][0] + got[i-1][1]
				sumCur := got[i][0] + got[i][1]
				if sumPrev > sumCur {
					t.Errorf("pairs not sorted by sum: pair[%d]=%d, pair[%d]=%d",
						i-1, sumPrev, i, sumCur)
				}
			}
		})
	}
}

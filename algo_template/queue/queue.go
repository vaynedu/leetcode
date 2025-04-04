package queue

type Queue []int

func (q *Queue) Push(x int) {
	*q = append(*q, x)
}
func (q *Queue) Pop() int {
	if len(*q) == 0 {
		return -1
	}
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}
func (q *Queue) Top() int {
	if len(*q) == 0 {
		return -1
	}
	return (*q)[0]
}

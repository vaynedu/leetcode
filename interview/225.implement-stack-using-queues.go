package interview

// 225. 用队列实现栈
// 单队列：每次 push 后把前面元素移到后面
type MyStack struct {
    queue []int
}

func StackConstructor() MyStack {
    return MyStack{}
}

func (this *MyStack) Push(x int) {
    this.queue = append(this.queue, x)
    // 把新元素之前的元素重新放到后面
    for i := 0; i < len(this.queue)-1; i++ {
        this.queue = append(this.queue, this.queue[0])
        this.queue = this.queue[1:]
    }
}

func (this *MyStack) Pop() int {
    if len(this.queue) == 0 {
        return 0
    }
    val := this.queue[0]
    this.queue = this.queue[1:]
    return val
}

func (this *MyStack) Top() int {
    if len(this.queue) == 0 {
        return 0
    }
    return this.queue[0]
}

func (this *MyStack) Empty() bool {
    return len(this.queue) == 0
}

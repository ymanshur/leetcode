type MyQueue struct {
    items []int
}


func Constructor() MyQueue {
    return MyQueue{items: []int{}}
}


func (this *MyQueue) Push(x int)  {
    this.items = append(this.items, x)
}


func (this *MyQueue) Pop() int {
    val := this.items[0]
    this.items[0] = 0
    this.items = this.items[1:]
    return val
}


func (this *MyQueue) Peek() int {
    return this.items[0]
}


func (this *MyQueue) Empty() bool {
    return len(this.items) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
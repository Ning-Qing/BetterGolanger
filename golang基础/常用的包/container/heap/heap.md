# heap
```go
type Interface interface {
	sort.Interface
	Push(x interface{}) // 将x添加为元素len()
	Pop() interface{}   // 删除并返回元素len()-1
}
// 初始化堆
heap.Init()
// 添加元素到堆上
heap.Push()
// 删除堆上最大或最小元素(取决于sort)并返回
heap.Pop()
```
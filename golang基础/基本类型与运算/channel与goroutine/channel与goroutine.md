# channel与goroutine
channel是Golang在语言层面提供的goroutine间的通信方式，与管道通信系统思想一致

## channel
[源码详见](./example/src/chan.go)
### channel结构
```go
type hchan struct {
	qcount   uint           //  当前队列中剩余元素个数
	dataqsiz uint           //  环形队列长度，即可以存放的元素个数
	buf      unsafe.Pointer //  环形队列指针
	elemsize uint16 // 每个元素的大小
	closed   uint32 // 标识关闭状态
	elemtype *_type // 元素类型
	sendx    uint   // 队列下标，指示元素写入时存放到队列中的位置
	recvx    uint   //  队列下标，指示元素从队列的该位置读出
	recvq    waitq  //  等待读消息的goroutine队列
	sendq    waitq  //  等待写消息的goroutine队列
	lock mutex // 互斥锁，chan不允许并发读写
}
```
- 由此可见channel数据结构由缓冲队列，元素类型信息，goroutine等待队列和互斥锁组成。
- 同时该结构也满足了管道机制必须提供的三方面协调能力，互斥、同步、确认对方是否存在

### 环形队列
chan内部实现了一个环形队列作为其缓冲区，队列的长度是创建chan时指定的
![](./static/环形队列.png)
- dataqsize: 表示队列长度，可缓存6个元素
- buf：指向队列内存
- sendx：表示后续写入数据的位置
- recvx: 表示从该位置读取数据

### goroutine等待队列
#### goroutine等待队列结构
```go
type waitq struct {
	first *sudog
	last  *sudog
}
```
从channel读数据，如果channel缓冲区为空或者没有缓冲区，当前goroutine会被阻塞。向channel写数据，如果channel缓冲区已满或者没有缓冲区，当前goroutine会被阻塞。
- 因读阻塞的goroutine会被向channel写入数据的goroutine唤醒
- 因写阻塞的goroutine会被从channel读数据的goroutine唤醒
#### sudog结构
```go
type sudog struct {
	g *g // 一个goroutine可以拥有多个sudog  

    // 队列指针 由此可见是双向链
	next *sudog 
	prev *sudog

	elem unsafe.Pointer //数据元素（可能指向堆栈）

	acquiretime int64
	releasetime int64
	ticket      uint32

    // isSelect 表示 g 正在参与一个选择
    // 所以 g.selectDone 必须是 CAS 才能赢得唤醒竞赛。
	isSelect bool
    // success 表示通过通道 c 的通信是否成功。
    // 如果 goroutine 因为值通过通道 c 传递而被唤醒，则为真
    // 如果因为 c 关闭而被唤醒，则为假。
	success bool

	parent   *sudog 
	waitlink *sudog 
	waittail *sudog 
	c        *hchan
}
```
#### 唤醒
```go
func (q *waitq) dequeue() *sudog {
	for {
        // 从等待队列出队一个sudog
		sgp := q.first
		if sgp == nil {
			return nil
		}
		y := sgp.next
		if y == nil {
			q.first = nil
			q.last = nil
		} else {
			y.prev = nil
			q.first = y
			sgp.next = nil 
		}
        // 因为sudog和g是多对一的关系
        // 所以可能当前的sudog对应的g已经被其他sudog唤醒
        // 所以因该跳过当前这个sudog
		if sgp.isSelect && !atomic.Cas(&sgp.g.selectDone, 0, 1) {
			continue
		}
		return sgp
	}
}
```
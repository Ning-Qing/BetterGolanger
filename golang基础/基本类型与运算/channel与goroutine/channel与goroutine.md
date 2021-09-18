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

## channal与goroutine
#### 向channel中写入数据
- 如果等待接收队列recvq不为空，说明缓冲区中没有数据或者没有缓冲区，此时直接从recvq取出G,并把数据写入，最后把该G唤醒，结束发送过程
- 如果缓冲区中有空余位置，将数据写入缓冲区，结束发送过程
- 如果缓冲区中没有空余位置，将待发送数据写入G，将当前G加入sendq，进入等待，等待被读goroutine唤醒
#### 从channel中读取数据
- 如果等待发送队列sendq不为空，且没有缓冲区，直接从sendq中取出G，把G中数据读出，最后把G唤醒，结束读取过程
- 如果等待发送队列sendq不为空，此时说明缓冲区已满，从缓冲区中首部读出数据，把G中数据写入缓冲区尾部，把G唤醒，结束读取过程
- 如果缓冲区中有数据，则从缓冲区取出数据，结束读取过程
- 将当前goroutine加入recvq，进入等待，等待被写goroutine唤醒
#### 一个十分重要的结构——sudog
```go
type sudog struct {
	g *g // 一个goroutine可以拥有多个sudog  

    // 队列指针 由此可见是双向链
	next *sudog 
	prev *sudog

	elem unsafe.Pointer //待读或被写的数据（可能指向堆栈)

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
sudog有三个重要部分，g、elem、isSelect和success
- g：sudog绑定的goroutine，一个goroutine绑定多个sudog
- elem：存储指向被写入或者待读数据的指针
- isSelect和success：因为有g有多个sudog，所以必须要能表示当前的状态参数 
#### 如何选取一个等待的goroutine
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
#### 如何唤醒一个等待的goroutine
```go
func send(c *hchan, sg *sudog, ep unsafe.Pointer, unlockf func(), skip int) {
	if raceenabled {
		if c.dataqsiz == 0 {
			racesync(c, sg)
		} else {
			racenotify(c, c.recvx, nil)
			racenotify(c, c.recvx, sg)
			c.recvx++
			if c.recvx == c.dataqsiz {
				c.recvx = 0
			}
			c.sendx = c.recvx // c.sendx = (c.sendx+1) % c.dataqsiz
		}
	}
	if sg.elem != nil {
		sendDirect(c.elemtype, sg, ep)
		sg.elem = nil
	}
	gp := sg.g
	// 解锁hchan  代表sudog完成了阻塞操作
	unlockf()
	// 将g.param指向完成阻塞操作的sudog 以唤醒goroutine
	gp.param = unsafe.Pointer(sg)
	sg.success = true
	if sg.releasetime != 0 {
		sg.releasetime = cputicks()
	}
	goready(gp, skip+1)
}
```
param 是指针字段,目前以三种方式使用：
1. 当一个channel操作唤醒一个阻塞的goroutine时，它设置param指向完成阻塞操作的sudog。在完成channel操作后，应设置为nil，进入等待状态
2. 通过gcAssistAlloc1向其调用者发信号通知goroutine完成了GC周期
3. 通过debugCallWrap向新的goroutine传递参数

#### 如何读写数据
```go
// 写数据
func sendDirect(t *_type, sg *sudog, src unsafe.Pointer) {
	// dst 从读等待队列中选取的g的数据区域
	dst := sg.elem
	// 写个屏障
	typeBitsBulkBarrier(t, uintptr(dst), uintptr(src), t.size)
	// 拷贝 src是写入的数据
	memmove(dst, src, t.size)
}
// 读数据 反过来就是了
func recvDirect(t *_type, sg *sudog, dst unsafe.Pointer) {
	src := sg.elem
	typeBitsBulkBarrier(t, uintptr(dst), uintptr(src), t.size)
	memmove(dst, src, t.size)
}
```
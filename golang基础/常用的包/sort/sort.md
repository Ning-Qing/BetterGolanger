# Sort

```go
package sort
type Interface interface {
    Len() int            // 获取元素数量
    Less(i, j int) bool // i，j是序列元素的下标。
    Swap(i, j int)        // 交换元素
}

// obj实现sort.Interface{}
sort.Sort(obj) 
```

package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	subAdd(t, 1, 2, 3)
	t.Run("subtest",func(t *testing.T) {
		subAdd(t,1,2,5)
	})	
}

func subAdd(t *testing.T, a, b, res int) bool {
	if v := Add(a, b); res != v {
		t.Fatalf("Add(%d + %d) return %d,not is %d", a, b, v, res)
		return false
	}
	return true
}

func BenchmarkAdd(b *testing.B) {
	if res := Add(1, 2); res != 3 {
		b.Fatalf("Add(1 + 2 ")
	}
}

// type BenchmarkResult struct {
//     N         int           // 迭代次数
//     T         time.Duration // 基准测试花费的时间
//     Bytes     int64         // 一次迭代处理的字节数
//     MemAllocs uint64        // 总的分配内存的次数
//     MemBytes  uint64        // 总的分配内存的字节数
// }

/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
type Heap struct {
	d []data
}

type data struct {
	key   int
	value int
}

func (h Heap) Len() int {
	return len(h.d)
}

func (h *Heap) Swap(i, j int) {
	h.d[i], h.d[j] = h.d[j], h.d[i]
}

func (h Heap) Less(i, j int) bool {
	return h.d[i].value < h.d[j].value
}

func (h *Heap) Pop() interface{} {
	back := h.d[len(h.d)-1]
	h.d = h.d[:len(h.d)-1]
	return back
}

func (h *Heap) Push(x interface{}) {
	h.d = append(h.d, x.(data))
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	h := &Heap{d:make([]data,0,k+1)}
	heap.Init(h)
	for key, value := range m {
		heap.Push(h,data{key: key, value: value})
		for h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int,0,k)
	for i:=0;i<k;i++{
		res = append(res, heap.Pop(h).(data).key)
	}
	return res
}
// @lc code=end


package main

import "fmt"

type Heap002 struct {
	Size int
	Elem []int
}

func newHeap002() *Heap002 {
	return &Heap002{Size: 0, Elem: []int{0}}
}

// Insert002 向堆中插入元素
func (h *Heap002) Insert002(x int)  {
	h.Elem = append(h.Elem, x)
	h.Size++
	if h.Size >= 1 {
		Swim002(h.Elem, h.Size)
	}
}

// delMax 堆中删除最大元素
func (h *Heap002) delMax() int {
	max := h.Elem[1]
	Swap002(h.Elem, 1, h.Size)
	h.Elem = h.Elem[:len(h.Elem)-1]
	h.Size--
	Sink002(h.Elem, 1, h.Size)
	return max
}


// 上浮
func Swim002(s []int, N int) {
	for {
		// 堆中root元素时 退出循环
		if N <= 1 {
			break
		}
		// 插入元素的父节点的索引
		k := N/2
		// 父节点处的元素大于插入元素则已经有序 退出循环
		if s[k] >= s[N] {
			break
		}
		// 交换N处元素和k处元素
		Swap002(s, k, N)
		// 进行递归向上判断
		N = N/2
	}
}

// 下沉
func Sink002(s []int, k, N int)  {
	for {
		i := 2*k
		if i > N {
			break
		}
		if i<N && s[i+1]>s[i] {
			i++
		}
		if s[k] > s[i] {
			break
		}
		Swap002(s, k, i)
		k = i
	}
}

// 交换元素
func Swap002(s []int, i, j int)  {
	s[i], s[j] = s[j], s[i]
}

// 堆排序
func HeapSort002(h Heap002) []int {
	heap := make([]int, 0)
	heap = append(heap, h.Elem...)
	size := h.Size
	sortHeap := make([]int, h.Size, h.Size)
	for i:=0; i<h.Size; i++ {
		m := heap[1]
		Swap002(heap, 1, size)
		heap = heap[:len(heap)-1]
		size--
		Sink002(heap, 1, size)
		sortHeap[i] = m
	}
	return sortHeap
}


func main() {
	h := newHeap002()
	// 向堆中插入元素
	s := []int{9, 0, 6, 5, 8, 2, 1, 7, 4, 3, 45, 12, 56, 78, 48, 100, 13, 24}
	for _, v := range s {
		h.Insert002(v)
	}
	// 打印有序后的堆
	fmt.Println("堆中的元素: ", h.Elem, "堆中元素个数: ", h.Size)
	// 进行堆排序
	d := HeapSort002(*h)
	fmt.Println(d)
	// 删除堆中最大的元素
	fmt.Println("删除的元素: ", h.delMax())
	fmt.Println("删除的元素: ", h.delMax())
	fmt.Println("删除的元素: ", h.delMax())
	fmt.Println("删除的元素: ", h.delMax())
	// 堆中剩下的元素
	fmt.Println("堆中的元素: ", h.Elem, "堆中元素个数: ", h.Size)
}


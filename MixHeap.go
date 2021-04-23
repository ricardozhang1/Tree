package main

import "fmt"

type Heap struct {
	Size int
	Elems []int
}

func NewHeap(maxSize int) *Heap {
	h := new(Heap)
	h.Elems = make([]int, maxSize, maxSize)
	return h
}

//func (h *Heap) Push(x int)  {
//	h.Size++
//
//	// i是要插入节点的下标
//	i := h.Size
//	for {
//		if i<= 0 {
//			break
//		}
//
//		// parent为父亲节点的下标
//		parent := (i-1)/2
//		// 如果父亲节点小于等于插入的值，则说明大小没有颠倒，可以退出
//		if h.Elems[parent] <= x {
//			break
//		}
//
//		// 互换当前父亲节点与要插入的值
//		h.Elems[i] = h.Elems[parent]
//		i = parent
//	}
//	h.Elems[i] = x
//}
//
//func (h *Heap) Pop() int {
//	if h.Size == 0 {
//		return 0
//	}
//
//	// 取出根节点
//	ret := h.Elems[0]
//
//	// 将最后一个节点的值提到根节点上
//	h.Size--
//	x := h.Elems[h.Size]
//
//	i := 0
//	for {
//		// a，b为左右两个子节点的下标
//		a := 2*i + 1
//		b := 2*i + 2
//		// 没有左子树
//		if a >= h.Size {
//			break
//		}
//
//		// 有右子树，找两个子节点中较小的值
//		if b < h.Size && h.Elems[b] < h.Elems[a] {
//			a = b
//		}
//
//		// 父亲小 直接退出
//		if h.Elems[a] >= x {
//			break
//		}
//		// 交换
//		h.Elems[i] = h.Elems[a]
//		i = a
//	}
//	h.Elems[i] = x
//	return ret
//}

func (h *Heap) Display()  {
	if h.Size > 0 {
		fmt.Printf("Size: %d, Elems: %#v\n", h.Size-1, h.Elems[1:h.Size])
	} else {
		fmt.Println("Size: 0, Elems: []int{}")
	}

}

//Push 向堆中插入元素
func (h *Heap) Push(x int) {
	// 使用堆排序来实现堆有序
	h.Size++
	i := h.Size  // 新元素插入的索引
	k := h.Size/2  // 新插入节点的父节点索引
	h.Elems[i] = x
	for k>0 {
		Sink(h.Elems, k, i)
		k--
	}
}

func (h *Heap) Pop() int {
	if h.Size == 0 {
		return -1
	}
	i := h.Size
	for i>0 {
		h.Elems[1] = h.Elems[i]
		//h.Elems[i] = 0
		i--
		//fmt.Println(h.Elems)
		Sink(h.Elems, 1, i)
	}

	return h.Elems[i+1]
}

// 下沉(由上至下的堆有序)
func Sink(s []int, k, N int) {
	for {
		i := 2*k
		if i>N {  // 保证该节点k 是非叶子节点
			break
		}
		if i<N && s[i+1] > s[i] {  // 选择较大的子节点 2k 2k+1
			i++
		}
		if s[k] >= s[i] {  // 没下沉到底就构造好堆了
			break
		}
		Swap(s, k, i)
		k = i
	}
}


func Swap(s []int, i, j int)  {
	s[i], s[j] = s[j], s[i]
}

func main() {
	h := NewHeap(10)
	h.Display()

	h.Push(3)
	h.Push(6)
	h.Push(7)
	h.Push(27)
	h.Push(1)
	h.Push(2)
	h.Push(3)
	h.Push(5)
	h.Push(9)
	h.Display()

	fmt.Println(h.Pop())
	//h.Display()
	fmt.Println(h.Pop())
	//h.Display()
	fmt.Println(h.Pop())
	//h.Display()
	fmt.Println(h.Pop())
	//h.Display()
	fmt.Println(h.Pop())
	//h.Display()

}




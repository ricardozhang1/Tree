package main

import "fmt"

/*
堆排序
s[0]不用，时间元素从角标1开始
父节点元素大于子节点元素
左子节点角标为2*k
右子节点角标为2*k+1
父节点角标为k/2
*/

func HeapSort(s []int) {
	N := len(s)-1  // 1-10
	// s[0]不用，实际元素数量和最后一个元素的角标都为N
	// 构造堆
	// 如果给两个已构造好的堆添加了一个共同的父节点，
	// 将新添加的节点做一次下沉将构造一个新堆，
	// 由于叶子节点都可以看做一个构造好的堆，所以
	// 可以从最后一个非叶子节点开始下沉，直至根节点，
	// 最后一个非叶子节点是最后一个叶子
	// 节点的父节点，角标为N/2

	for k:=N/2; k>=1; k-- {  // 进行构造堆
		sink(s, k, N)
	}
	fmt.Println(s)

	// 下沉排序
	for N>1 {
		swap(s, 1, N)  // 将大的放在数组后面，升序排序
		N--
		sink(s, 1, N)
	}
}

// 下沉(由上至下的堆有序)
func sink(s []int, k, N int) {
	for {
		i := 2*k
		if i>N {  // 保证该节点k 是非叶子节点
			break
		}
		if i<N && s[i+1] > s[i] {  // 有两个节点 选择较大的子节点 2k 2k+1
			i++
		}
		if s[k] >= s[i] {  // 没下沉到底就构造好堆了
			break
		}
		swap(s, k, i)
		k = i
	}
}

func swap(s []int, i, j int)  {
	s[i], s[j] = s[j], s[i]
}

// 测试
func main() {
	s := []int{-1,9, 0, 6, 5, 8, 2, 1, 7, 4, 3}
	//fmt.Println(s[1:])
	HeapSort(s)
	fmt.Println(s)
}


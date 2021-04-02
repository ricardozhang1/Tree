package main

import "fmt"

// Binary Search Trees
// 二叉查找树
type BinaryTree struct {
	root *node
	n int
}

// create the node
func newNode(k, v int) *node {
	return &node{k: k, v: v, sz: 1}
}

// create binary tree
func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

// 增加或修改
func (b *BinaryTree) Put(k, v int) {
	b.root, _ = put(b.root, k, v)
}

// 查找
func (b *BinaryTree) Get(k int) int {
	return get(b.root, k)
}

// 树的大小
func (b *BinaryTree) Size() int {
	return size(b.root)
}

// 选出最小键
func (b *BinaryTree) Min() int {
	return min(b.root).k
}

// 删除最小键
func (b *BinaryTree) DeleteMin()  {
	b.root = deleteMin(b.root)
}

// 删除
func (b *BinaryTree) Delete(k int)  {
	b.root = delete(b.root, k)
}

// 遍历
// 中序遍历
func (b *BinaryTree) MidOrder() {
	midOrder(b.root)
}

// 前序遍历
func (b *BinaryTree) PreOrder() {
	preOrder(b.root)
}

// 后序遍历
func (b *BinaryTree) LastOrder() {
	lastOrder(b.root)
}


// node
type node struct {
	k, v, sz int  // 键、值、大小
	left, right *node  // 左右子节点
}

// 在以nd为根节点的树下增加或修改一个节点
// 如果创建了新的节点，第二个参数返回true，
// 如果只是修改，第二个参数返回false
func put(nd *node, k, v int) (*node, bool) {
	if nd == nil {
		return newNode(k, v), true
	}
	hasNew := false

	if k < nd.k {
		nd.left, hasNew = put(nd.left, k, v)
	} else if k > nd.k {
		nd.right, hasNew = put(nd.right, k, v)
	} else {
		nd.v = v  // 仅修改，不会增加节点，就不更新树的大小
	}

	if hasNew {
		updateSize(nd)  // 如果创建了新节点就更新树的大小
	}

	return nd, hasNew
}

// 在以nd为根节点的树中获取键为k的值
func get(nd *node, k int) int {
	if nd == nil {
		return -1
	}

	if k < nd.k {
		return get(nd.left, k)
	} else if k > nd.k {
		return get(nd.right, k)
	} else {
		return nd.v
	}
}

// 获取以nd为根节点的树的大小
func size(nd *node) int {
	if nd == nil {
		return 0
	}
	return nd.sz
}

// 更新以nd为根节点的树的大小
func updateSize(nd *node) {
	if nd == nil {
		return
	}
	nd.sz = size(nd.left) + size(nd.right) + 1
}

// 选出以nd为根节点的树的最小键节点
func min(nd *node) *node {
	if nd == nil {
		return nil
	}
	if nd.left != nil {
		return min(nd.left)
	}
	return nd
}

// 删除以nd为根节点的树的最小键节点
// 返回被删除的节点
func deleteMin(nd *node) *node {
	if nd == nil {
		return nil
	}

	if nd.left == nil {  // 找到最小节点
		nd = nd.right  // 用右子节点代替自己
	} else {  // 还有更小的
		nd.left = deleteMin(nd.left)
	}
	updateSize(nd)
	return nd
}

// 删除以nd为根节点的树并且键为k的节点
func delete(nd *node, k int) *node {
	if nd == nil {
		return nil
	}

	if k < nd.k {
		nd.left = delete(nd.left, k)
	} else if k > nd.k {
		nd.right = delete(nd.right, k)
	} else {
		// 删除的的非叶子节点
		if nd.right == nil {
			return nd.left
		}
		if nd.left == nil {
			return nd.right
		}
		// 同时具有两个子节点
		// 先找出大于本节点的最小节点作为后续节点
		t := nd
		nd.k = min(t.right).k
		// 删除
		deleteMin(t.right)
		// 用后续节点代替本节点
		nd.left = t.left
	}
	updateSize(nd)
	return nd
}

// 以nd为根节点的中序遍历
func midOrder(nd *node) {
	if nd == nil {
		return
	}
	// 先打印左子节点
	midOrder(nd.left)
	// 按照次序打印根节点
	fmt.Println(nd.k)
	// 打印右子树
	midOrder(nd.right)
}
// 以nd为根节点的前序遍历
func preOrder(nd *node)  {
	if nd == nil {
		return
	}
	// 先打印根节点
	fmt.Println(nd.k)
	//然后打印左子节点
	preOrder(nd.left)
	//最后打印右子节点
	preOrder(nd.right)
}

// 以nd为根节点的前序遍历
func lastOrder(nd *node) {
	if nd == nil {
		return
	}
	// 先遍历右子节点
	lastOrder(nd.right)
	// 然后遍历左子节点
	lastOrder(nd.left)
	// 最后遍历根节点
	fmt.Println(nd.k)
}


func main() {
	b := NewBinaryTree()
	b.Put(5, 5555)
	b.Put(4, 4444)
	b.Put(7, 7777)
	b.Put(6, 6666)
	b.Put(8, 8888)

	//fmt.Println(b.Get(8))
	//fmt.Println(b.Size())
	//fmt.Println(b.Min())
	//b.DeleteMin()
	//fmt.Println(b.Get(4))
	//b.Delete(7)
	//fmt.Println(b.Size())
	//b.MidOrder()

	//b.PreOrder()
	b.LastOrder()
}









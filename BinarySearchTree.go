package main

import "fmt"

type BinaryTree002 struct {
	root *Node002
	size int
}

// newBinaryTree002 创建一个二叉树
func newBinaryTree002() *BinaryTree002 {
	return &BinaryTree002{}
}

// Put 往二叉树中插入元素节点
func (b *BinaryTree002) Put(k int, v interface{})  {
	b.root, _ = Put(b.root, k, v)
}

// Get 获取树中的节点
func (b *BinaryTree002) Get(k int) interface{} {
	return Get(b.root, k)
}

// Size 获取树种节点数目
func (b *BinaryTree002) Size() int {
	return sizeNode(b.root)
}

// Min 获取树中的最小节点
func (b *BinaryTree002) Min() interface{} {
	return Min(b.root).value
}

// DeleteMin 删除树中最小的元素
func (b *BinaryTree002) DeleteMin()  {
	b.root = DeleteMin(b.root)
}

// Delete 删除树中指定节点
func (b *BinaryTree002) Delete(k int)  {
	b.root = Delete(b.root, k)
}

// MidOrder 中序遍历
func (b *BinaryTree002) MidOrder()  {
	MidOrder(b.root)
}

// PreOrder 前序遍历
func (b *BinaryTree002) PreOrder()  {
	PreOrder(b.root)
}

// LastOrder 后序遍历
func (b *BinaryTree002) LastOrder()  {
	LastOrder(b.root)
}

// LayerOrder 层序遍历
func (b *BinaryTree002) LayerOrder()  {
	LayerOrder(b.root)
}

// 节点结构
type Node002 struct {
	key int
	value interface{}
	sz int
	left *Node002
	right *Node002
}

// newNode002 创建新的节点
func newNode002(k int, v interface{}) *Node002 {
	return &Node002{key: k, value: v, sz: 1}
}

func Put(n *Node002, k int, v interface{}) (*Node002, bool) {
	if n == nil {
		return newNode002(k, v), true
	}
	addNode := false
	if k < n.key {
		n.left, addNode = Put(n.left, k, v)
	} else if k > n.key {
		n.right, addNode = Put(n.right, k, v)
	} else {
		n.value = v
	}
	if addNode {
		// 修改节点sz数 非叶子节点才会调用
		updateNodeSize(n)
	}
	return n, addNode
}

func updateNodeSize(n *Node002) {
	// 参与到递归中
	if n == nil {
		return
	}
	n.sz = sizeNode(n.left) + sizeNode(n.right) + 1
}

func sizeNode(n *Node002) int {
	if n == nil {
		return 0
	}
	return n.sz
}

func Get(n *Node002, k int) interface{} {
	if n == nil {
		return -1
	}
	if k < n.key {
		return Get(n.left, k)
	} else if k > n.key {
		return Get(n.right, k)
	} else {
		return n.value
	}
}

func Min(n *Node002) *Node002 {
	if n == nil {
		return nil
	}
	if n.left != nil {
		return Min(n.left)
	}
	return n
}

func DeleteMin(n *Node002) *Node002 {
	if n == nil {
		return nil
	}
	if n.left == nil {
		n = n.right
	} else {
		n.left = DeleteMin(n.left)
	}
	updateNodeSize(n)
	return n

}

func Delete(n *Node002, k int) *Node002 {
	if n == nil {
		return nil
	}
	if k < n.key {
		n.left = Delete(n.left, k)
	} else if k > n.key {
		n.right = Delete(n.right, k)
	} else {
		// 删除非叶子节点
		if n.right == nil {
			return n.left
		}
		if n.left == nil {
			return n.left
		}
		// 同时具有两个节点
		// 先找出大于本节点的最小节点作为后续节点
		t := n
		n.key = Min(t.right).key
		// 删除
		DeleteMin(t.right)
		// 用后续节点代替本节点
		n = t.left
	}
	UpdateSize(n)
	return n
}

func UpdateSize(n *Node002) {
	if n == nil {
		return
	}
	n.sz = Size(n.left) + Size(n.right) + 1
}

func Size(n *Node002) int {
	if n == nil {
		return 0
	}
	return n.sz
}

func MidOrder(n *Node002) {
	if n == nil {
		return
	}
	MidOrder(n.left)
	fmt.Printf("%v\n", n.value)
	MidOrder(n.right)
}

func PreOrder(n *Node002) {
	if n == nil {
		return
	}
	fmt.Printf("%v\n", n.value)
	PreOrder(n.left)
	PreOrder(n.right)
}

func LastOrder(n *Node002) {
	if n == nil {
		return
	}
	LastOrder(n.right)
	LastOrder(n.left)
	fmt.Printf("%v\n", n.value)
}

func LayerOrder(n *Node002) {
	// 需要创建一个队列
	temp := make([]*Node002, 0)
	temp = append(temp, n)
	var nd *Node002
	for len(temp) > 0 {
		nd = temp[0]
		fmt.Println(nd.value)
		if nd.left != nil {
			temp = append(temp, nd.left)
		}
		if nd.right != nil {
			temp = append(temp, nd.right)
		}
		temp = temp[1:]
	}
}

func main() {
	b := newBinaryTree002()
	b.Put(5, "王五")
	b.Put(3, "张三")
	b.Put(4, "李四")
	b.Put(6, "赵六")
	//b.Put(2, "陈二")
	//fmt.Println(b.root.sz)
	//fmt.Println(b.Get(4))
	//fmt.Println(b.Min())

	//b.DeleteMin()
	//fmt.Println(b.Get(4))
	//fmt.Println(b.root.sz)
	fmt.Println("===========层序遍历==============")
	b.LayerOrder()
	fmt.Println("===========前序遍历==============")
	b.PreOrder()
	fmt.Println("===========中序遍历==============")
	b.MidOrder()
	fmt.Println("===========后序遍历==============")
	b.LastOrder()
}


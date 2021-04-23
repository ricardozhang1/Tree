package main

import "fmt"

// 定义颜色
const (
	RED = true
	BLACK = false
)

// 左倾红黑树
type LLRBTree struct {
	root *LLRBTNode  // 树的根节点
}

// 左倾红黑树节点
type LLRBTNode struct {
	value int  // 值
	times int  // 值出现的次数
	left *LLRBTNode  // 左子树
	right *LLRBTNode  // 右子树
	color bool  // 父亲指向该节点的链接颜色
}

// 新建一课空树
func NewLLRBTree() *LLRBTree {
	return &LLRBTree{}
}

// 对节点实现左旋转
func RotateLeft(h *LLRBTNode) *LLRBTNode {
	if h == nil {
		return nil
	}
	x := h.right
	h.right = x.left
	x.left = h
	x.color = h.color
	h.color = RED
	return x
}

// 对节点实现最旋转
func RotateRight(h *LLRBTNode) *LLRBTNode {
	if h == nil {
		return nil
	}
	x := h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = RED
	return x
}

// 颜色转换
func ColorChange(h *LLRBTNode)  {
	if h == nil {
		return
	}

	h.color = !h.color
	h.left.color = !h.left.color
	h.right.color = !h.right.color
}

// 左倾红黑树添加元素
func (tree *LLRBTree) Add(value int)  {
	// 从根节点开始添加元素，因为可能调整，所以需要将返回的节点赋值回根节点
	tree.root = tree.root.Add(value)
	// 根节点的链接永远都是黑色的
	tree.root.color = BLACK
}

func (node *LLRBTNode) Add(v int) *LLRBTNode {
	// 插入的节点为空，将其链接颜色设置为红色，并返回
	if node == nil {
		return &LLRBTNode{
			value: v,
			color: RED,
		}
	}

	// 插入的元素重复
	if v == node.value {
		node.times = node.times + 1
	} else if v > node.value {
		// 插入的元素比节点值大，往右子树插入
		node.right = node.right.Add(v)
	} else {
		// 插入的元素比节点值小，往左子节点插入
		node.left = node.left.Add(v)
	}

	// 辅助变量
	nowNode := node
	// 右链接为红色，那么要进行左旋，确保树是左倾的
	// 这里完成后就可以结束了，因为插入操作，新插入的右红链接左旋后，
	//nowNode节点不会出现连续两个红左链接，因为它只有一个左红链接
	if IsRed(nowNode.right) && !IsRed(nowNode.left) {
		nowNode = RotateLeft(nowNode)
	} else {
		// 连续两个左连接为红色，那么进行右旋
		if IsRed(node.left) && IsRed(node.left.left) {
			nowNode = RotateRight(nowNode)
		}

		// 旋转后，可能左右链接都为红色，需要变色
		if IsRed(node.left) && IsRed(nowNode.right) {
			ColorChange(nowNode)
		}
	}
	return nowNode
}

// 判断节点颜色
func IsRed(node *LLRBTNode) bool {
	if node == nil {
		return false
	}
	return node.color == RED
}

// 红色左移
// 节点h是红节点，其左儿子和左孙子都是黑节点，
// 左移后使得其左儿子或左儿子的左儿子有一个是红色节点
func MoveRedLeft(h *LLRBTNode) *LLRBTNode {
	// 应该确保isRed(h) && isRed(h.left) && isRed(h.left.left)
	ColorChange(h)
	// 右儿子有左红链接
	if IsRed(h.right.left) {
		// 对又儿子右旋
		h.right = RotateRight(h.right)
		// 再左旋
		h = RotateLeft(h)
		ColorChange(h)
	}
	return h
}

// 红色右移
// 节点 h 是红节点，其右儿子和右儿子的左儿子都为黑节点，右移后使得其右儿子或右儿子的右儿子有一个是红色节点
func MoveRedRight(h *LLRBTNode) *LLRBTNode {
	// 应该确保isRed(h) && !isRed(h.right) && !idRed(h.right.left)
	ColorChange(h)
	// 左儿子有左红连接
	if IsRed(h.left.left) {
		// 右旋
		h = RotateLeft(h)
		// 变色
		ColorChange(h)
	}
	return h
}

// 左倾红黑树删除元素
func (tree *LLRBTree) Delete(v int)  {
	// 当前找不到值时直接返回
	if tree.Find(v) == nil {
		return
	}

	if !IsRed(tree.root.left) && !IsRed(tree.root.right) {
		// 左右子树都是黑节点，那么先将根节点变为红色节点，方便后面的红色左移或右移
		tree.root.color = RED
	}
	tree.root = tree.root.Delete(v)

	// 最后，如果根节点非空，永远要为黑节点，赋值黑色
	if tree.root != nil {
		tree.root.color = BLACK
	}
}

// 对该节点所在的子树删除元素
func (node *LLRBTNode) Delete(v int) *LLRBTNode {
	// 辅助变量
	nowNode := node
	// 删除的元素比子树根节点小，需要从左子树删除
	if v < nowNode.value {
		// 因为从左子树删除，所以需要判断是否要红色左移
		if !IsRed(nowNode.left) && !IsRed(nowNode.left.left) {
			// 左儿子和左儿子的儿子都不是红色节点，那么没法递归下去，先红色左移
			nowNode = MoveRedLeft(nowNode)
		}
		// 现在可以从左子树中删除了
		nowNode.left = nowNode.left.Delete(v)
	} else {
		// 删除的元素等于或大于树的根节点
		// 左节点为红色，那么需要又旋，方便后面可以红色右移
		if IsRed(nowNode.left) {
			nowNode = RotateRight(nowNode)
		}

		// 值相等，且没有右孩子节点，那么该节点一定是要被删除的叶子节点，直接删除
		// 为什么呢，反证，她没有右儿子，但有左儿子，
		// 以为左倾红黑树的特征，那么左儿子一定是红色，
		// 但是前面的语句已经把红色左儿子右旋到右边，不应该出现右儿子为空。
		if v == nowNode.value && nowNode.right == nil {
			return nil
		}

		// 因为从右子树删除，所以要判断是否需要红色右移
		if !IsRed(nowNode.right) && !IsRed(nowNode.right.left) {
			// 右儿子和右儿子的左儿子都不是红色节点，那么没法递归下去，先红色右移
			nowNode = MoveRedRight(nowNode)
		}

		// 删除的节点找到了，她是中间的节点，需要用最小后驱动节点来替换它，然后删除最小后驱节点
		if v == nowNode.value {
			minNode := nowNode.right.FindMinValue()
			nowNode.value = minNode.value
			nowNode.times = minNode.times

			// 删除其最小节点
			nowNode.right = nowNode.right.DeleteMin()
		} else {
			// 删除的元素比子树的根节点大，需要从右子树删除
			nowNode.right = nowNode.right.Delete(v)
		}
	}
	return nowNode.FixUp()
}

// 查找指定节点
func (tree *LLRBTree) Find(v int) *LLRBTNode {
	if tree.root == nil {
		// 如果树是空得
		return nil
	}
	return tree.root.Find(v)
}

func (node *LLRBTNode) Find(v int) *LLRBTNode {
	if v == node.value {
		// 如果该节点刚刚等于该值，那么返回该节点
		return node
	} else if v < node.value {
		// 如果查找的值小于节点值，从节点的左子树开始找
		if node.left == nil {
			// 左子树为空，表示找不到该值了，返回nil
			return nil
		}
		return node.left.Find(v)
	} else {
		// 如果查找的值大于节点值，从节点的右子树开始找
		if node.right == nil {
			// 右子树为空，表示找不到该值了，返回nil
			return nil
		}
		return node.right.Find(v)
	}
}

// 对该节点所在的子树删除最小元素
func (node *LLRBTNode) DeleteMin() *LLRBTNode {
	// 辅助变量
	nowNode := node
	// 没有左子树，那么删除它自己
	if nowNode.left == nil {
		return nil
	}

	// 判断是否需要红色左移，因为最小元素的左子树中
	if !IsRed(nowNode.left) && !IsRed(nowNode.left.left) {
		nowNode = MoveRedLeft(nowNode)
	}

	// 递归从左子树删除
	nowNode.left = nowNode.left.DeleteMin()

	// 修复左倾红黑树的特征
	return nowNode.FixUp()
}

// 修复左倾红黑树特征
func (node *LLRBTNode) FixUp() *LLRBTNode {
	// 辅助变量
	nowNode := node

	// 红链接在左边，左旋转恢复，让红链接只出现在左边
	if IsRed(nowNode.right) {
		nowNode = RotateLeft(nowNode)
	}

	// 连续两个左连接为红色，那么进行右旋
	if IsRed(nowNode.left) && IsRed(nowNode.left.left) {
		nowNode = RotateRight(nowNode)
	}

	// 旋转后，可能左右链接都为红色，需要变色
	if IsRed(nowNode.left) && IsRed(nowNode.right) {
		ColorChange(nowNode)
	}
	return nowNode
}

// 找出最小值的节点
func (tree *LLRBTree) FindMinValue() *LLRBTNode {
	if tree.root == nil {
		// 如果是空树，返回空
		return nil
	}
	return tree.root.FindMinValue()
}

func (node *LLRBTNode) FindMinValue() *LLRBTNode {
	// 左子树为空，表面已经是最左的节点了，该值就是最小值
	if node.left == nil {
		return node
	}

	// 一直左子树递归
	return node.left.FindMinValue()
}

// 找出最大值节点
func (tree *LLRBTree) FindMaxValue() *LLRBTNode {
	if tree.root == nil {
		return nil
	}
	return tree.root.FindMaxValue()
}


func (node *LLRBTNode) FindMaxValue() *LLRBTNode {
	// 右子树为空，表面已经是最右的节点，该值就是最大值
	if node.right == nil {
		return node
	}
	return node.right.FindMaxValue()
}

// 中序遍历
func (tree *LLRBTree) MidOrder()  {
	tree.root.MidOrder()
}

func (node *LLRBTNode) MidOrder()  {
	if node == nil {
		return
	}
	// 先打印左子树
	node.left.MidOrder()
	// 然后打印中间节点
	for i:=0; i<=int(node.times); i++ {
		fmt.Println(node.value)
	}
	//最后打印右子节点
	node.right.MidOrder()
}

// 验证是不是棵左倾红黑树
func (tree *LLRBTree) IsLLRBTree() bool {
	if tree == nil || tree.root == nil {
		return true
	}

	// 判断是否是棵二分查找树
	if !tree.root.IsBST() {
		return false
	}

	// 判断该树是否遵循2-3树，也就是红链接只能在左边，不能连续有两个红链接
	if !tree.root.Is23() {
		return false
	}

	// 判断树是否平衡，也就是任意一个节点到叶子节点，经过的黑色链接数量相同
	// 计算根节点到最左边叶子节点的黑链接数量
	blackNum := 0
	x := tree.root
	for x != nil {
		if !IsRed(x) {  // 黑色链接
			blackNum = blackNum + 1
		}
		x = x.left
	}

	if !tree.root.IsBalanced(blackNum) {
		return false
	}
	return true
}

// 节点所在的子树是否是一棵二分查找树
func (node *LLRBTNode) IsBST() bool {
	if node == nil {
		return true
	}

	// 左子树非空，那么根节点必须大于左儿子节点
	if node.left != nil {
		if node.value > node.left.value {

		} else {
			fmt.Printf("father: %#v, lchild: %#v, rchild: %#v\n", node.value, node.left.value, node.right.value)
			return false
		}
	}

	// 右子树非空，那么根节点必须小于右儿子节点
	if node.right != nil {
		if node.value < node.right.value {

		} else {
			fmt.Printf("father: %#v, lchild: %#v, rchild: %#v\n", node, node.left, node.right)
			return false
		}
	}

	// 左子树也要判断是否是平衡查找树
	if !node.left.IsBST() {
		return false
	}

	// 右子树也要判断是否是平衡查找树
	if !node.right.IsBST() {
		return false
	}
	return true
 }

// 节点所在的子树是否遵循2-3树
func (node *LLRBTNode) Is23() bool {
	if node == nil {
		return true
	}

	// 不允许右倾红链接
	if IsRed(node.right) {
		fmt.Printf("father: %#v, rchild: %#v\n", node, node.right)
		return false
	}

	// 不允许连续两个左红链接
	if IsRed(node) && IsRed(node.left) {
		fmt.Printf("father: %#v, lchild: %#v\n", node, node.left)
		return false
	}

	// 左子树也要判断是否遵循2-3树
	if !node.left.Is23() {
		return false
	}

	// 右子树也要判断是否遵循2-3树
	if !node.right.Is23() {
		return false
	}
	return true
}

// 节点所在的子树是否平衡，是否有blackNum个很链接
func (node *LLRBTNode) IsBalanced(blackNum int) bool {
	if node == nil {
		return blackNum==0
	}

	if !IsRed(node) {
		blackNum = blackNum - 1
	}

	if !node.left.IsBalanced(blackNum) {
		fmt.Println("node.left to leaf black link is not.", blackNum)
		return false
	}

	if !node.right.IsBalanced(blackNum) {
		fmt.Println("node.right to leaf black link is not.", blackNum)
		return false
	}
	return true
}



func main() {
	tree := NewLLRBTree()
	values := []int{2, 3, 7, 10, 10, 10, 10, 23, 9, 102, 109, 111, 112, 113}
	for _, v := range values {
		tree.Add(v)
	}

	// 找到最大值或最小值的节点
	fmt.Println("find min value:", tree.FindMinValue())
	fmt.Println("find max value:", tree.FindMaxValue())

	// 查找不存在的99
	node := tree.Find(99)
	if node != nil {
		fmt.Println("find it 99!")
	} else {
		fmt.Println("not find it 99!")
	}

	// 查找存在的9
	node = tree.Find(9)
	if node != nil {
		fmt.Println("find it 9!")
	} else {
		fmt.Println("not find it 9!")
	}

	tree.MidOrder()

	// 删除存在的9后，再查找9
	//tree.Delete(9)
	//tree.Delete(10)
	//tree.Delete(2)
	//tree.Delete(3)
	//tree.Add(4)
	//tree.Add(3)
	//tree.Add(10)
	//tree.Delete(111)
	//node = tree.Find(9)
	if node != nil {
		fmt.Println("find it 9!")
	} else {
		fmt.Println("not find it 9!")
	}

	if tree.IsLLRBTree() {
		fmt.Println("is a llrb tree")
	} else {
		fmt.Println("is not llrb tree")
	}
}







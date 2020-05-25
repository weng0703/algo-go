package tree

type BST struct {
	*BinaryTree
	compareFunc func(v, nodeV interface{}) int
}

func NewBST(rootV interface{}, compareFunc func(v, nodeV interface{}) int) *BST {
	if nil == compareFunc {
		return nil
	}
	return &BST{BinaryTree: NewBinaryTree(rootV), compareFunc: compareFunc}
}

//查找
func (this *BST) Find(v interface{}) *Node {
	p := this.root
	if nil != p {
		complexResult := this.compareFunc(v, p.data)
		if complexResult == 0 {
			return p
		} else if complexResult > 0 {
			p = p.right
		} else {
			p = p.left
		}
	}
	return nil
}

func (this *BST) Insert(v interface{}) bool {
	p := this.root
	for nil != p {
		compareResult := this.compareFunc(v, p.data)
		if compareResult == 0 {
			return false
		} else if compareResult > 0 {
			if nil == p.right {
				p.right = NewNode(v)
				break
			}
			p = p.right
		} else {
			if nil == p.left {
				p.left = NewNode(v)
				break
			}
			p = p.left
		}
	}
	return true
}

func (this *BST) Delete(v interface{}) bool {
	var pp *Node = nil //被删除节点的父节点
	p := this.root
	deleteLeft := false
	for nil != p {
		compareResult := this.compareFunc(v, p.data)
		if compareResult > 0 {
			pp = p
			p = p.right
			deleteLeft = false
		} else if compareResult < 0 {
			pp = p
			p = p.left
			deleteLeft = true
		} else {
			break
		}
	}

	if nil == p { //需要删除的节点不存在
		return false
	} else if nil == p.left && nil == p.right { //删除的是一个叶子节点
		if nil != pp {
			if deleteLeft {
				pp.left = nil
			} else {
				pp.right = nil
			}
		} else { //根节点
			this.root = nil
		}
	} else if nil != p.right { //删除的是一个有右孩子，不一定有左孩子的节点
		//找到p节点右孩子的最小节点
		pq := p      //当前节点
		q := p.right //当前节点的右节点
		fromRight := true
		for nil != q.left { //向左走到底
			pq = q
			q = q.left
			fromRight = false
		}
		if fromRight { //只有一个右节点
			pq.right = nil
		} else {
			pq.left = nil
		}
		q.left = p.left
		q.right = p.right
		if nil == pp { //根节点被删除
			this.root = q
		} else {
			if deleteLeft {
				pp.left = q
			} else {
				pp.right = q
			}
		}
	} else { //删除的是一个只有左孩子的节点
		if nil != pp {
			if deleteLeft {
				pp.left = p.left
			} else {
				pp.right = p.left
			}
		} else {
			if deleteLeft {
				this.root = p.left
			} else {
				this.root = p.left
			}
		}
	}

	return true

}

//查找最小值
func (this *BST) Min() *Node {
	p := this.root
	if this.root.left != nil {
		p = p.left
	}
	return p
}

//查找最大值
func (this *BST) Max() *Node {
	p := this.root
	if this.root.right != nil {
		p = p.right
	}
	return p
}

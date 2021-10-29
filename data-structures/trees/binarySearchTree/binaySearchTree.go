package binarySearchTree

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func (b *BST) insert(data int) {
	if b.root == nil {
		b.root = &Node{data, nil, nil}
		return
	}
	b.insertNode(b.root, data)
}

func (b *BST) insertNode(node *Node, data int) {
	if data < node.data {
		if node.left == nil {
			node.left = &Node{data, nil, nil}
			return
		}
		b.insertNode(node.left, data)
	} else {
		if node.right == nil {
			node.right = &Node{data, nil, nil}
			return
		}
		b.insertNode(node.right, data)
	}
}

func (b *BST) inOrder() {
	if b.root == nil {
		return
	}
	b.inOrderNode(b.root)
}

func (b *BST) inOrderNode(node *Node) {
	if node == nil {
		return
	}
	b.inOrderNode(node.left)
	fmt.Println(node.data)
	b.inOrderNode(node.right)
}

func (b *BST) postOrder() {
	if b.root == nil {
		return
	}
	b.postOrderNode(b.root)
}

func (b *BST) postOrderNode(node *Node) {
	if node == nil {
		return
	}
	b.postOrderNode(node.left)
	b.postOrderNode(node.right)
	fmt.Println(node.data)
}

func (b *BST) preOrder() {
	if b.root == nil {
		return
	}
	b.preOrderNode(b.root)
}

func (b *BST) preOrderNode(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.data)
	b.preOrderNode(node.left)
	b.preOrderNode(node.right)
}

func (b *BST) search(data int) bool {
	if b.root == nil {
		return false
	}
	return b.searchNode(b.root, data)
}

func (b *BST) searchNode(node *Node, data int) bool {
	if node == nil {
		return false
	}
	if data == node.data {
		return true
	}
	if data < node.data {
		return b.searchNode(node.left, data)
	}
	return b.searchNode(node.right, data)
}

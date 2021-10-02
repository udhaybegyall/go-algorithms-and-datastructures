package linkedList

import "fmt"

type Node struct {
	value int
	next  *Node
}

type linkedList struct {
	head *Node
}

// Add value at the begning of linked-list.
func (ll *linkedList) addAtBegning(value int) {

	var node = Node{
		value: value,
		next:  &Node{},
	}

	node.value = value

	if node.next != nil {
		node.next = ll.head
	}
	ll.head = &node
}

// Add value at the end of linked-list.
func (ll *linkedList) addAtEnd(value int) {
	var node = &Node{
		value: value,
		next:  &Node{},
	}

	node.value = value
	node.next = nil

	var lastNode *Node
	lastNode = ll.getLastNode()

	if lastNode == nil {
		lastNode.next = node
	}
}

// Add value after given number.
func (ll *linkedList) addAfter(node_to_be_searched int, value int) {
	var node = &Node{
		value: value,
		next:  &Node{},
	}

	node.value = value
	node.next = nil

	var nodeLocation *Node
	nodeLocation = ll.searchNode(node_to_be_searched)

	if nodeLocation == nil {
		fmt.Println("Insertion not possible")
		return
	} else {
		node.next = nodeLocation.next
		nodeLocation.next = node
	}
}

// Returns lastnode position
func (ll *linkedList) getLastNode() *Node {

	var node *Node
	var lastNode *Node

	for node = ll.head; node != nil; node = node.next {
		if node.next == nil {
			lastNode = node
		}
	}
	return lastNode
}

// Returns the location of the node.
func (ll *linkedList) searchNode(value int) *Node {

	var node *Node

	for node = ll.head; node != nil; node = node.next {
		if node.value == value {
			return node
		}
	}
	return nil
}

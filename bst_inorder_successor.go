package main

// using methods from bst.go

func findInOrderSuccessor(root *BSTNode, data int) *BSTNode{
	curr := findNode(root, data)
	if curr.Right != nil {
		return findMin(curr.Right)
	} else {
		var a1 *BSTNode
		var a2 *BSTNode = root
		for a2 != curr {
			if curr.Data < a2.Data {
				a1 = a2
				a2 = a2.Left
			} else {
				a2 = a2.Right
			}
		}
		return a1
	}
}

func main() {

}

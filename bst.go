package main

type BSTNode struct {
	Data int
	Left *BSTNode
	Right *BSTNode
}

func findNode(root *BSTNode, data int) *BSTNode {
	if root == nil {
		return nil
	}
	if root.Data == data {
		return root
	} else if data > root.Data {
		return findNode(root.Right, data)
	} else {
		return findNode(root.Left, data)
	}
}

func findNodeItr(root *BSTNode, data int) *BSTNode {
	for root != nil && root.Data != data {
		if data > root.Data {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return root
}

func findMin(root *BSTNode) *BSTNode {
	min := root
	for root.Left != nil {
		min = root.Left
	}
	return min
}

func findMax(root *BSTNode) *BSTNode {
	max := root
	for root.Right != nil {
		max = root.Right
	}
	return max
}

func height(root *BSTNode) int {
	if root == nil {
		return -1
	}
	lHeight := height(root.Left)
	rHeight := height(root.Right)
	if lHeight > rHeight {
		return 1 + lHeight
	} else {
		return 1 + rHeight
	}
}

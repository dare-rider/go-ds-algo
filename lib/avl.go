package main

type Node struct{
    Val int
    Left *Node
    Right *Node
    Height int
    Size int
}

func AvlInsert(root *Node, val int) *Node {
    if root == nil {
        node := Node{Val: val, Height: 0, Size: 1}
        return &node
    }
    if val < root.Val {
        root.Left = AvlInsert(root.Left, val)
    } else {
        root.Right = AvlInsert(root.Right, val)
    }
    root.Size += 1
    balance := root.Left.Height - root.Right.Height
    if balance > 1 {
        lsubH := -1
        rsubH := -1
        if root.Left.Left != nil {
           lsubH = root.Left.Left.Height
        }
        if root.Left.Right != nil {
           rsubH = root.Left.Right.Height
        }
        if lsubH >= rsubH {
            // LL
            return rotateRight(root)
        } else {
            // LR
            root.Left = rotateLeft(root.Left)
            return rotateRight(root)
        }
    } else if balance < -1 {
        lsubH := -1
        rsubH := -1
        if root.Right.Left != nil {
           lsubH = root.Right.Left.Height
        }
        if root.Right.Right != nil {
           rsubH = root.Right.Right.Height
        }
        if lsubH >= rsubH {
            // RL
            root.Right = rotateRight(root.Right)
            return rotateLeft(root)
        } else {
            // RR
            return rotateLeft(root)
        }
    }
    h := 0
    if root.Left != nil {
        h = root.Left.Height
    }
    if root.Right != nil && root.Right.Height > h {
        h = root.Right.Height
    }
    root.Height = 1 + h
    return root
}

func rotateLeft(node *Node) *Node {

}

func rotateRight(node *Node) *Node {

}

package main

import (
	"fmt"
	"strconv"
)

//TreeNode, a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//根据前序和后序遍历构造二叉树
func constructFromPrePost(pre []int, post []int) *TreeNode {
	if len(pre) == 0 || len(pre) != len(post) {
		return nil
	}
	root := &TreeNode{
		Val: pre[0],
	}
	size := len(post)
	if size == 1 {
		return root
	}
	for i := 0; i < size; i++ {
		if post[i] == pre[1] {
			root.Left = constructFromPrePost(pre[1:i+2], post[:i+1])
			root.Right = constructFromPrePost(pre[i+2:], post[i+1:size-1])
			break
		}
	}
	return root
}

// pre search
func (node *TreeNode) traverse() {
	if node == nil {
		return
	}
	fmt.Print(strconv.Itoa(node.Val) + " ")
	node.Left.traverse()
	node.Right.traverse()
}

func (root *TreeNode) BFS() [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) != 0 {
		cl := []int{}
		size := len(queue)
		for i := 0; i < size; i++ {
			current := queue[0]
			queue = queue[1:]
			cl = append(cl, current.Val)
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		res = append(res, cl)
	}
	return res
}

func (root *TreeNode) DFS(depth int, res *[][]int) {
	if root == nil {
		return
	}
	for len(*res) <= depth {
		*res = append(*res, []int{})
	}
	root.Left.DFS(depth+1, res)
	root.Right.DFS(depth+1, res)
	(*res)[depth] = append((*res)[depth], root.Val)
}

func main() {
	var pre = []int{1, 2, 4, 5, 3, 6, 7}
	post := []int{4, 5, 2, 6, 7, 3, 1}
	var tree = constructFromPrePost(pre, post)
	tree.traverse()
	var bfslist = tree.BFS()
	fmt.Println(bfslist)
	var dfslist = [][]int{}
	tree.DFS(0, &dfslist)
	fmt.Println(dfslist)
}

package main

import (
	"fmt"
	"strconv"
)

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	r := x
	for r > x/r {
		r = (r + x/r) / 2
	}
	return int(r)
}

func perfectnum(n int) bool {
	if n == 0 {
		return true
	} else {
		var sum int = 0
		for i := 1; i*i <= n; i++ {

			if n%i == 0 {
				sum += i
				if i*i != n {
					sum += (n / i)
				}
			}

		}
		return sum-n-n == 0
	}
}

//TreeNode, a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

func (node *TreeNode) traverse() {
	if node == nil {
		return
	}
	node.Left.traverse()
	fmt.Print(strconv.Itoa(node.Val) + " ")
	node.Right.traverse()
}

func main() {
	fmt.Println(mySqrt(8))
	fmt.Println(perfectnum(0))
	var pre = []int{1, 2, 4, 5, 3, 6, 7}
	var post = []int{4, 5, 2, 6, 7, 3, 1}
	var tree = constructFromPrePost(pre, post)
	tree.traverse()
}

package tree

import (
	"fmt"
	"strings"
)

func binaryTreePaths(root *TreeNode) []string {
	var cur []string
	var solutions [][]string
	binaryTreePathsRecur(root, &cur, &solutions)

	var result []string
	for _, solu := range solutions {
		result = append(result, strings.Join(solu, "->"))
	}

	return result
}

func binaryTreePathsRecur(node *TreeNode, cur *[]string, solutions *[][]string) {
	if node == nil {
		return
	}

	*cur = append(*cur, fmt.Sprintf("%d", node.Val))
	defer func() {
		*cur = (*cur)[:len(*cur)-1]
	}()

	if node.Left == nil && node.Right == nil {
		*solutions = append(*solutions, append([]string{}, *cur...))
		return
	}

	binaryTreePathsRecur(node.Left, cur, solutions)
	binaryTreePathsRecur(node.Right, cur, solutions)
}

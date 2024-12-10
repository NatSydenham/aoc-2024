package day7

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/NatSydenham/aoc-2024/lib/file"
	"github.com/NatSydenham/aoc-2024/lib/slice"
)

type Tree struct {
	Root *Node
}

type Node struct {
	Value    int
	Children []*Node
}

func (tree *Tree) HasPath(target int, start int) bool {
	var dfs func(node *Node) bool
	dfs = func(node *Node) bool {

		if node == nil {
			return false
		}

		if len(node.Children) == 0 {
			return node.Value == target
		}

		for _, child := range node.Children {
			if dfs(child) {
				return true
			}
		}

		return false
	}

	return dfs(tree.Root)
}

func (tree *Tree) BuildTree(start int, nums []int) {
	var build func(curr int, idx int) *Node
	build = func(curr, idx int) *Node {
		if idx >= len(nums) {
			return &Node{Value: curr}
		}

		node := &Node{Value: curr}

		add := curr + nums[idx]
		mul := curr * nums[idx]

		node.Children = append(node.Children, build(add, idx+1))
		node.Children = append(node.Children, build(mul, idx+1))

		return node
	}

	tree.Root = build(start, 1)
}

func (tree *Tree) BuildTreeWithConcat(start int, nums []int) {
	var build func(curr int, idx int) *Node
	build = func(curr, idx int) *Node {
		if idx >= len(nums) {
			return &Node{Value: curr}
		}

		node := &Node{Value: curr}

		add := curr + nums[idx]
		mul := curr * nums[idx]
		conc, _ := strconv.Atoi(strconv.Itoa(curr) + strconv.Itoa(nums[idx]))

		node.Children = append(node.Children, build(add, idx+1))
		node.Children = append(node.Children, build(mul, idx+1))
		node.Children = append(node.Children, build(conc, idx+1))

		return node
	}

	tree.Root = build(start, 1)
}

func ExecutePart1() {
	start := time.Now()
	lines := file.Readlines("./data/day7.txt")
	sum := 0

	for _, line := range lines {
		sections := strings.Split(line, ": ")
		target, _ := strconv.Atoi(sections[0])

		nums := slice.Map(strings.Split(sections[1], " "), func(item string) int {
			result, _ := strconv.Atoi(item)
			return result
		})

		tree := &Tree{}
		tree.BuildTree(nums[0], nums)

		if tree.HasPath(target, nums[0]) {
			sum += target
		}
	}

	fmt.Println(sum, "|", time.Since(start))
}

func ExecutePart2() {
	start := time.Now()
	lines := file.Readlines("./data/day7.txt")
	sum := 0

	for _, line := range lines {
		sections := strings.Split(line, ": ")
		target, _ := strconv.Atoi(sections[0])

		nums := slice.Map(strings.Split(sections[1], " "), func(item string) int {
			result, _ := strconv.Atoi(item)
			return result
		})

		tree := &Tree{}
		tree.BuildTreeWithConcat(nums[0], nums)

		if tree.HasPath(target, nums[0]) {
			sum += target
		}
	}

	fmt.Println(sum, "|", time.Since(start))
}

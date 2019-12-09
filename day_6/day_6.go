package main

import (
	"fmt"

	"github.com/grmpflh27/aoc_2019/aoc2019_shared"
)

type orbit struct {
	name     string
	parent   *orbit
	children []*orbit
}

func (o *orbit) totalOrbits() int {
	cnt := 0
	cur := o
	for {
		cur = cur.parent
		if cur == nil {
			break
		}
		cnt++
	}
	return cnt
}

func (o *orbit) parents() []*orbit {
	var allParents []*orbit
	curParent := o.parent
	for curParent != nil {
		allParents = append(allParents, curParent)
		curParent = curParent.parent
	}
	return allParents
}

func build(input [][]string) *orbit {
	root := &orbit{
		"COM",
		nil,
		nil,
	}
	add(input, root)
	return root
}

func add(input [][]string, node *orbit) {
	addChildren(input, node)
	for _, c := range node.children {
		add(input, c)
	}
}

func addChildren(input [][]string, parent *orbit) {
	var children []*orbit
	for _, fromTo := range input {
		if fromTo[0] == parent.name {
			curOrbit := &orbit{
				fromTo[1],
				parent,
				nil,
			}
			children = append(children, curOrbit)
		}
	}
	parent.children = children
}

func dfsTraverse(node *orbit) []*orbit {
	stack := []*orbit{node}
	var path []*orbit

	for len(stack) > 0 {
		// poor mans pop
		vertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if contains(path, vertex) {
			continue
		}
		path = append(path, vertex)
		for _, c := range vertex.children {
			stack = append(stack, c)
		}
	}
	return path
}

func contains(path []*orbit, node *orbit) bool {
	for _, o := range path {
		if o.name == node.name {
			return true
		}
	}
	return false
}

func find(node *orbit, name string) (*orbit, error) {
	path := dfsTraverse(node)
	for _, p := range path {
		if p.name == name {
			return p, nil
		}
	}
	return nil, fmt.Errorf("Did not find %v in tree", name)
}

func getCommonParent(firstParents []*orbit, secondParents []*orbit) *orbit {
	var commonParent *orbit
	if len(firstParents) > len(secondParents) {
		for _, node := range firstParents {
			if contains(secondParents, node) {
				commonParent = node
				break
			}
		}
	} else {
		for _, node := range secondParents {
			if contains(firstParents, node) {
				commonParent = node
				break
			}
		}
	}
	return commonParent
}

func main() {
	var day = 6
	fmt.Println("==========")
	fmt.Println("Day ", day)
	fmt.Println("==========")

	input := aoc2019_shared.LoadStr(day, ")")
	root := build(input)

	// 1
	totalOrbitCnt := 0
	fullPath := dfsTraverse(root)
	for _, node := range fullPath {
		curCnt := node.totalOrbits()
		//fmt.Printf("Node %v has %v total orbits\n", node.name, curCnt)
		totalOrbitCnt += curCnt
	}

	fmt.Printf("Answer 1: %v total orbits\n", totalOrbitCnt)

	// 2
	// find YOU and SANTA
	youNode, _ := find(root, "YOU")
	santaNode, _ := find(root, "SAN")

	// find common parent
	youParents := youNode.parents()
	santaParents := santaNode.parents()
	commonParent := getCommonParent(youParents, santaParents)

	// calculating distance to common parent (exclusive of YOU|SANTA)
	yDistance := youNode.parent.totalOrbits() - commonParent.totalOrbits()
	sDistance := santaNode.parent.totalOrbits() - commonParent.totalOrbits()

	fmt.Printf("Answer2: %v distance\n", yDistance+sDistance)
}

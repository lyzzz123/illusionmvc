package handler

import (
	"fmt"
	"strings"
)

type PathTreeNode struct {
	SubPath        string
	Children       map[string]*PathTreeNode
	Parent         *PathTreeNode
	HandlerWrapper *Wrapper
}

type PathTreeMap struct {
	Root *PathTreeNode
}

func (pathTreeMap *PathTreeMap) GetHandlerWrapper(path string) *Wrapper {
	subPaths := strings.Split(path, "/")
	findNode := pathTreeMap.Root
	for i := 1; i < len(subPaths); i++ {
		node, ok := findNode.Children[subPaths[i]]
		if !ok {
			nodeStar, okStar := findNode.Children["*"]
			if !okStar {
				return nil
			} else {
				findNode = nodeStar
			}
		} else {
			findNode = node
		}
	}
	return findNode.HandlerWrapper
}

func (pathTreeMap *PathTreeMap) PutHandlerWrapper(path string, handlerWrapper *Wrapper) {
	if pathTreeMap.Root == nil {
		pathTreeMap.Root = &PathTreeNode{
			SubPath:  "",
			Children: make(map[string]*PathTreeNode),
		}
	}
	subPaths := strings.Split(path, "/")
	findNode := pathTreeMap.Root
	for i := 1; i < len(subPaths); i++ {
		sp := subPaths[i]
		pathTreeNode, ok := findNode.Children[sp]
		if ok {
			findNode = pathTreeNode
		} else {
			child := &PathTreeNode{
				SubPath:  sp,
				Children: make(map[string]*PathTreeNode),
				Parent:   findNode,
			}
			findNode.Children[sp] = child
			findNode = child
		}
	}
	if findNode.HandlerWrapper != nil {
		panic("路径重复")
	}
	findNode.HandlerWrapper = handlerWrapper
}

func main() {
	pathTreeMap := &PathTreeMap{}
	pathTreeMap.PutHandlerWrapper("/a/*/c/*/fd", nil)
	pathTreeMap.PutHandlerWrapper("/a/*/d", nil)
	pathTreeMap.PutHandlerWrapper("/a/b/d/dd", nil)
	pathTreeMap.PutHandlerWrapper("/a/*/d", nil)
	fmt.Println("asdfasdf")

	fmt.Println(pathTreeMap.GetHandlerWrapper("/a/gsdf/c/gggggg/fd"))
	fmt.Println(pathTreeMap.GetHandlerWrapper("/a/g/d"))

}

package utils

import "reflect"
import (
	. "github.com/nzlov/go/array"
)

type TreeNode struct {
	parent   *TreeNode
	Data     interface{}
	Children *Array
	leaf     bool
}

func NewTreeData() *TreeNode {
	return NewTreeDataByData(nil)
}

func NewTreeDataByData(data interface{}) *TreeNode {
	return &TreeNode{nil, data, NewArray(), true}
}

func (this *TreeNode) AddChildren(node *TreeNode) {
	this.Children.Add(node)
	this.leaf = false
	node.SetParent(this)
}
func (this *TreeNode) AddChildrens(nodes ...*TreeNode) {
	for _, v := range nodes {
		this.AddChildren(v)
	}
}
func (this *TreeNode) Insert(index int, node *TreeNode) bool {
	if index > this.Children.Size() {
		return false
	}
	this.Children.Insert(index, node)
	node.parent = this
	this.leaf = false
	return true
}
func (this *TreeNode) Remove(node *TreeNode) bool {
	if this.Children.RemoveValue(node) {
		if this.Children.Size() == 0 {
			this.leaf = true
		}
		return true
	}
	return false
}
func (this *TreeNode) RemoveByIndex(index int) (*TreeNode, error) {
	n, err := this.Children.RemoveIndex(index)

	if this.Children.Size() == 0 {
		this.leaf = true
	}

	return n.(*TreeNode), err
}

func (this *TreeNode) Parent() *TreeNode {
	return this.parent
}

func (this *TreeNode) SetParent(p *TreeNode) {
	if this.parent != nil {
		this.Parent().Remove(this)
	}
	this.parent = p
}

type Tree struct {
	Root    *TreeNode
	Tag     string
	allNode map[string]*TreeNode
}

func NewTree() *Tree {
	return NewTreeByTag("Id")
}
func NewTreeByRootNode(root *TreeNode) *Tree {
	return NewTreeByRootNodeAndTag(root, "id")
}
func NewTreeByTag(tag string) *Tree {
	return NewTreeByRootNodeAndTag(nil, tag)
}
func NewTreeByRootNodeAndTag(root *TreeNode, tag string) *Tree {
	return &Tree{root, tag, make(map[string]*TreeNode)}
}
func (this *Tree) Add(parent, children *TreeNode) {
	parent.AddChildren(children)
	this.allNode[this.getTag(children)] = children
}
func (this *Tree) Insert(parent, children *TreeNode, index int) bool {
	return parent.Insert(index, children)
}
func (this *Tree) Remove(parent, node *TreeNode) {
	if parent.Remove(node) {
		delete(this.allNode, this.getTag(node))
	}
}
func (this *Tree) RemoveByIndex(parent *TreeNode, index int) (*TreeNode, error) {
	n, err := parent.RemoveByIndex(index)
	if err == nil {
		delete(this.allNode, this.getTag(n))
	}
	return n, err
}
func (this *Tree) AllNode() map[string]*TreeNode {
	return this.allNode
}

func (this *Tree) getTag(node *TreeNode) string {
	return reflect.ValueOf(node.Data).Elem().FieldByName(this.Tag).String()
}

func (this *Tree) GetTreeNode(n string) *TreeNode {
	return this.allNode[n]
}

func (this *Tree) PathToRoot(node *TreeNode) []*TreeNode {
	return this.pathToRootByDepth(node, 0)
}

func (this *Tree) pathToRootByDepth(node *TreeNode, depth int) []*TreeNode {
	var retNodes []*TreeNode

	if node == nil {
		if depth == 0 {
			return nil
		}
		retNodes = make([]*TreeNode, depth)
	} else {
		depth = depth + 1
		if node == this.Root {
			retNodes = make([]*TreeNode, depth)
		} else {
			retNodes = this.pathToRootByDepth(node.Parent(), depth)
		}
		retNodes[len(retNodes)-depth] = node
	}
	return retNodes
}

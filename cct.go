package gocct

import (
	"reflect"
)

type Tree interface {
	// Returns the root node of the tree
	Root() Node
	// Adds a child to the provided parent.
	// Returns true, nil if added successfully,
	// false, current parent if
	// the node already exists in the tree
	Add(parent Node, child Node) (success bool, otherParent Node)
	// Get the parent of the provided node
	Parent(child Node) (parent Node)
	// Depth first walk of the tree calling the provided function for each node visited
	Walk(func(depth int, node Node))

	// Remove(node Node) (bool)
	// GoWalk() chan<- Visit

	// Children of parent
	Children(parent Node) []Node
	//
}

type Visit struct {
	N Node
	I int
}

// Node, an interface{} type store in the tree.
// Tider then writing interface{} everywhere
type Node interface{}

// Fluent API for building a tree
type Builder interface {
	Add(n Node) Builder
	Down() Builder
	Up() Builder
	Build() Tree
}

// Create a new tree with a hidden root and return the root and the tree
func NewTree() (Tree, Node) {
	t := &ctree{
		make(map[Node][]Node),
		make(map[Node]Node),
		new(interface{}),
		true,
	}
	t.c2[t.root] = nil
	return t, t.root
}

// Creates a tree that allow non pointer node.
// Client beware, do not mutate object or pass non keyable objects
func NewTree_MutableNodes() (Tree, Node) {
	t := &ctree{
		make(map[Node][]Node),
		make(map[Node]Node),
		new(interface{}),
		false,
	}
	t.c2[(Node)(t.root)] = nil
	return (Tree)(t), t.root
}

// The internal structure of the centralised tree
type ctree struct {
	// Map of parent to slice of children
	p2 map[Node][]Node
	// Map of child to parent
	c2 map[Node]Node
	// Hidden root node of the tree
	root    Node
	ptrOnly bool
}

// Add a child to a parent and create the p2c and c2p pointers
// If the child already exists anywhere in the tree the operations fails and the
// existing parent is returned with the boolean status
func (t *ctree) Add(p Node, c Node) (bool, Node) {
	if t.ptrOnly {
		if reflect.TypeOf(c).Kind() != reflect.Ptr {
			panic("Child node is not a pointer")
		}
		if reflect.TypeOf(p).Kind() != reflect.Ptr {
			panic("Parent node is not a pointer")
		}
	}
	if orginal, exist := t.c2[c]; exist {
		return false, orginal
	}
	t.p2[p] = append(t.p2[p], c)
	t.c2[c] = p
	return true, nil
}

func (t *ctree) Parent(c Node) Node {
	return t.c2[c]
}

func (t *ctree) Children(p Node) []Node {
	return t.p2[p]
}

// A depth first walk of the tree calling the provided function at each node
func (t *ctree) Walk(f func(int, Node)) {
	for _, o := range t.p2[t.root] {
		walk(t, o, 0, f)
	}
}

func walk(t *ctree, node Node, depth int, f func(int, Node)) {
	f(depth, node)
	for _, o := range t.p2[node] {
		walk(t, o, depth+1, f)
	}
}

// Return the root of the tree
func (t *ctree) Root() Node {
	return t.root
}

// Builder's state
type builder struct {
	tree Tree
	curr Node
	last Node
}

func BuildTree() Builder {
	t, r := NewTree()
	b := &builder{t, r, nil}
	return b
}
func BuildTree_MutableNodes() Builder {
	t, r := NewTree_MutableNodes()
	b := &builder{t, r, nil}
	return b
}
func (b *builder) Add(n Node) Builder {
	ok, _ := b.tree.Add(b.curr, n)
	if ok {
		b.last = n
	}
	return b
}
func (b *builder) Down() Builder {
	b.curr = b.last
	return b
}
func (b *builder) Up() Builder {
	b.curr = b.tree.Parent(b.curr)
	b.last = b.curr
	return b
}
func (b *builder) Build() Tree {
	return b.tree
}

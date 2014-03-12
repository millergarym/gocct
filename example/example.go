package main

import (
	"fmt"
	"github.com/millergarym/gocct"
	"strings"
)

// Testing structure
type mys struct {
	s string
}

func main() {
	ct, r := gocct.NewTree()
	a := mys{"a"}
	ct.Add(r, &a)
	b := mys{"b"}
	c := mys{"c"}
	ct.Add(&a, &b)
	ct.Add(&a, &c)
	ct.Add(&a, &c) // ignored as c already in the tree
	x := 1
	ct.Add(r, &x)
	ct.Walk(p)

	fmt.Println()

	d := "d"
	d2 := "d"
	slice := &[]int{1, 2, 3, 4}
	gocct.BuildTree_MutableNodes().
		Add(mys{"a"}).Down().
		/**/ Add(&mys{"b"}).
		/**/ Add(mys{"c"}).
		/**/ Add(&d2).
		/**/ Add(d).
		/**/ Add("d").Up().
		Add(slice).
		Build().Walk(p)
}

// Function given to the walker for printing nodes
func p(d int, n gocct.Node) {
	s := strings.Repeat(" ", d)
	//	fmt.Println( reflect.TypeOf( n ) )
	switch v := n.(type) {
	case *mys:
		fmt.Printf("1)%s%v\n", s, v.s)
	case mys:
		fmt.Printf("2)%s%v\n", s, v.s)
	case *string:
		fmt.Printf("3)%s%v\n", s, v)
	case string:
		fmt.Printf("4)%s%v\n", s, v)
	case *int:
		fmt.Printf("5)%s%v\n", s, *v)
	case *[]int:
		fmt.Printf("6)%s%v\n", s, *v)
	default:
		fmt.Printf("*)%s%v - %T\n", s, n, n)
	}
}

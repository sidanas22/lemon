// package main
//
// import (
// 	"fmt"
// 	"math"
// )
//
// type Abser interface {
// 	Abs() float64
// }
//
// func main() {
// 	var a Abser
// 	f := MyFloat(-math.Sqrt2)
// 	v := Vertex{3, 4}
//
// 	// a = f
//
// 	a = &v
//
// 	fmt.Println(a.Abs())
// }
//
// type MyFloat float64
//
// func (f MyFloat) Abs() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }
//
// type Vertex struct {
// 	X, Y float64
// }
//
// func (v *Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.Y + v.Y*v.Y)
// }

// package main
//
// import "fmt"
//
// type I interface {
// 	M()
// }
//
// type T struct {
// 	S string
// }
//
// func (t T) M() {
// 	fmt.Println(t.S)
// }
//
// func main() {
// 	var i I = T{S: "Hello"}
// 	i.M()
// }

/* package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%+v, %T)\n", i, i)
} */

/**
 * interface-values-with-nil.go
**/

/* package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}

	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{S: "Hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%+v, %T)\n", i, i)
} */

/**
 * nil-interface-values.go
**/

/* package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
} */

/**
 * empty-interface.go
 *
 * empty interfaces may hold values of any type
 * they are mainly used for handling values of unknown types
**/

package main

import "fmt"

// type I interface{}

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

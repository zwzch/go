package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type Vertex1 struct {
	Lat, Long float64
}

var m = map[string]Vertex1{
	"Bell Labs": Vertex1{
		40.68433, -74.39967,
	},
	"Google": Vertex1{
		37.42202, -122.08408,
	},
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X + v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	fmt.Println(v1, p, v2, v3)

	a := make([]int, 5)

	fmt.Println(a)

	var b []int
	fmt.Println(b)
	b = append(b, 0)
	fmt.Println(b)
	b = append(b, 1)
	fmt.Println(b)
	b = append(b, 2, 3, 4, 5)
	fmt.Println(b)

	for i := range b {
		fmt.Println(i)
		b[i] = 1 << uint(i)
	}

	fmt.Printf("bis %v", b)

	for _, value := range b {
		fmt.Println(value)
	}

	fmt.Println(m)

	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("the value is", m["Answer"])

	m["Answer"] = 48
	fmt.Println("the value is", m["Answer"])

	v1, ok1 := m["Answer"]
	fmt.Println("value is", v1, "present", ok1)

	delete(m, "Answer")
	fmt.Println("the value is", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("value is", v, "present", ok)

	value := &Vertex{3, 4}
	fmt.Println(value.Abs())
	fmt.Println(value)

	value.Scale(5)
	fmt.Println(value.Abs())
	fmt.Println(value)

}

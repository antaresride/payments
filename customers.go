package main

import (
	"fmt"
)

type People struct {
	//lint:ignore U1000 will use in next PR
	age int
	//lint:ignore U1000 will use in next PR
	name string
}

type Merchant struct {
	name string //lint:ignore U1000 will use in next PR
}

type Buyer interface {
	buy() float64
	claim() string
	swap() string
}

func initial_customers() []string {
	return []string{"Luie", "Marlon", "David"}
}

// to do: associate customers and their ages
// hint map
func initial_customers_age(a []string, b []int32) []int32 {

	for i := range len(a) {
		fmt.Println(i)
	}
	return []int32{1, 2, 3, 4}

}

//lint:ignore U1000 will use in next PR
func (Merchant) buy() float64 {
	return 0.0
}

//lint:ignore U1000 will use in next PR
func (People) buy() float64 { //lint:ignore U1000 will use in next PR
	return 0.0
}

func (Merchant) M(p string) string {
	return p
}

func (*Merchant) N(p string) string {
	return p
}

func deposit() float64 {
	return Add(3.1, 4.3)
}

/*func main() {
	fmt.Println("Hi")
	p1 := People{}
	p2 := People{age: 13, name: "John"}
	fmt.Println(p1)
	fmt.Println(p2)
	m1 := Merchant{}
	fmt.Println(m1)
	fmt.Println("fisrt name in m1: ", m1.M("hi"))
	fmt.Println(m1.N("hello"))
	mp := &m1
	mp.name = "Gled"
	//fmt.Println(*mp)
	fmt.Println("new name in m1: ", mp.name)
	fmt.Println(m1)
	fmt.Println(*mp == m1)
}
*/
//method and interface

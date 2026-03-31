package main

import (
	"fmt"

	"rsc.io/quote"
)

type Payment struct {
	who string
}
type Provider interface {
	blockchain() string
}

func Hello() string {
	return "Hello!"
}
func (Payment) Provider() string {
	return "Genpay!"
}

func (Payment) Crack() string {
	return "Hi"
}

func main() {
	// 1. Fixed "ineffectual assignment" by using the values or removing them
	a := 3
	b := 5

	var amount [10]float64
	amount[0] = 12.00

	if a < b {
		fmt.Println("Hi")
	} else {
		fmt.Println("Hey")
	}
	fmt.Println(a)
	fmt.Println(b)

	var payment1 = Payment{who: "Gled"}
	var payment2 = Payment{}

	fmt.Println(payment1)
	fmt.Println(payment2)

	// 2. Fixed "identical expressions" (don't compare a variable to itself)
	fmt.Println(payment1 == payment2)

	fmt.Println(Hello())
	fmt.Println(payment1.Provider())
	fmt.Println(payment2.Provider())

	fmt.Println(amount[0])

	// 3. Fixed "func value not called": quote.Go is a function, call it with ()
	fmt.Println(quote.Go())
	// For test the push with Zed container
	fmt.Println(payment1.Crack())

	p1 := People{}

	fmt.Println(p1)

	fmt.Println(Add(3.4, 5.5))

	fmt.Println(deposit())

	fmt.Println(initial_customers())

	ages := []int32{1, 2, 3, 4}

	fmt.Println(initial_customers_age(initial_customers(), ages))

	map1 := map[string]int64{
		"1": 2,
		"2": 4,
	}
	fmt.Println(len(map1))

}

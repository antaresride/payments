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

func main() {
	var a = 1
	a = 2
	a = 3
	b := 4
	b = 5

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

	fmt.Println(payment1 == payment1)

	fmt.Println(Hello())

	fmt.Println(payment1.Provider())

	fmt.Println(payment2.Provider())

	fmt.Println(amount[0])

	fmt.Println(quote.Go)
}

package main

import "fmt"

func showType[T any](variable T) {
	fmt.Printf("%v , %T", variable, variable)
}

func main() {

	x := "James bond"
	y := 42
	z := true

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	showType(x)
	showType(y)
	showType(z)

}

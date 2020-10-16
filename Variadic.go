package main

import "fmt"

//sum Regresa el valor más grande del slice
func sum(args ...int) int{
	mayor := 0
	for i := range args{
		if args[i]>mayor{
			mayor = args[i]
		}
	}
	return mayor
}

func main() {
	a := []int{4,2,7,1,4,3,9,10,543}
	fmt.Println(sum(a...))

}
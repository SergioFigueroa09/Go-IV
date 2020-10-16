package main
import "fmt"

//generadorImpares genera impares empezando desde 1
func generadorImpares() func() int{
	i := 1

	return func() int{
		var impar = i
		i += 2
		return impar
	}
	
}

func main(){
	nextImpar := generadorImpares()
	fmt.Println(nextImpar())
	fmt.Println(nextImpar())
	fmt.Println(nextImpar())
}
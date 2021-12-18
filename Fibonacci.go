package main

func GetNfibonacciElements(num1 int64, num2 int64, numberOfElements int64) []int64 {

	lista := make([]int64, 2)
	lista[0] = num1
	lista[1] = num2

	var i, j int64 = 0, numberOfElements - 2

	for i < j {
		lista = append(lista, (lista[len(lista)-1] + lista[len(lista)-2]))
		i++
	}
	return lista
}

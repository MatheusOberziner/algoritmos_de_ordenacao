package main

import "fmt"

// Função para mesclar dois subarrays ordenados
func merge(arr []int, left []int, right []int) {
	i, j, k := 0, 0, 0
	nL, nR := len(left), len(right)

	// Mesclar os elementos em ordem crescente
	for i < nL && j < nR {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	// Adicionar elementos restantes de left[] (se houver)
	for i < nL {
		arr[k] = left[i]
		i++
		k++
	}

	// Adicionar elementos restantes de right[] (se houver)
	for j < nR {
		arr[k] = right[j]
		j++
		k++
	}
}

// Função principal de merge sort
func mergeSort(arr []int) {
	n := len(arr)

	// Caso base: se o tamanho do array for menor ou igual a 1, retorna
	if n <= 1 {
		return
	}

	// Divide o array em duas metades
	mid := n / 2
	left := make([]int, mid)
	right := make([]int, n-mid)

	// Preenche as metades esquerda e direita
	for i := 0; i < mid; i++ {
		left[i] = arr[i]
	}
	for i := mid; i < n; i++ {
		right[i-mid] = arr[i]
	}

	// Chamadas recursivas para as duas metades
	mergeSort(left)
	mergeSort(right)

	// Mescla as duas metades ordenadas
	merge(arr, left, right)
}

func main() {
	// Array de exemplo
	arr := []int{12, 11, 13, 5, 6, 7}

	fmt.Println("Array original:", arr)

	// Chamada para ordenar o array usando Merge Sort
	mergeSort(arr)

	fmt.Println("Array ordenado:", arr)
}

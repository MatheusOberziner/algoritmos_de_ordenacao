package main

import (
	"fmt"
)

// Função que particiona o array ao redor do pivô
func partition(arr []int, low, high int) int {
	// Escolhe o último elemento como pivô
	pivot := arr[high]
	// Inicializa o índice do menor elemento
	i := low - 1

	// Percorre o array, movendo os elementos menores que o pivô para a esquerda
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Coloca o pivô na posição correta, para que todos os elementos menores
	// estejam à esquerda, e os maiores à direita
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// Implementa o algoritmo de ordenação Quick Sort
func quickSort(arr []int, low, high int) {
	if low < high {
		// Particiona o array e obtém o índice do pivô
		pivot := partition(arr, low, high)

		// Chama recursivamente quickSort para as duas metades do array
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}

func main() {
	// Array de exemplo
	arr := []int{5, 9, 3, 7, 2, 8, 6, 1, 4}

	fmt.Println("Array original:", arr)

	// Chamada para ordenar o array usando Quick Sort
	quickSort(arr, 0, len(arr)-1)

	fmt.Println("Array ordenado:", arr)
}

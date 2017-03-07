package main

import "fmt"

func reverseSweep(numbers []int) {

		var firstIndex int = len(numbers) - 1
		var secondIndex int = len(numbers) - 2

		for secondIndex >= 0 {
			var firstNumber int = numbers[firstIndex]
      var secondNumber int = numbers[secondIndex]

      if firstNumber > secondNumber {
        numbers[firstIndex] = secondNumber
        numbers[secondIndex] = firstNumber
      }

			firstIndex--
			secondIndex--
		}

}

func reverseBubbleSort(numbers []int) {
  var N int = len(numbers)
  var i int
  for i = 0; i < N; i++ {
    reverseSweep(numbers)
  }
}

func main() {

	var numbers []int = []int{21, 4, 2, 13, 10, 0, 19, 11, 7, 5, 23, 18, 9, 14, 6, 8, 1, 20, 17, 3, 16, 22, 24, 15, 12}
  fmt.Println("Unsorted:", numbers)
  reverseBubbleSort(numbers)
  fmt.Println("Sorted:  ", numbers)

}

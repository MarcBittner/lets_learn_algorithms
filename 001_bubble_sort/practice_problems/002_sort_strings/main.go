package main

import "fmt"
import "strings"

func greater(a, b string) bool {
	if strings.ToLower(a) > strings.ToLower(b) {
		return true
	} else {
		return false
	}
}


func sweep(animals []string) {

//	var numberOfElements = len(animals)
		var firstIndex int = 0
		var secondIndex int = 1

// 		fmt.Println("numberOfElements:", len(animals))
//		fmt.Println("firstIndex:", firstIndex)
//		fmt.Println("secondIndex:", secondIndex)
// 		fmt.Println("animals value #1:", animals[firstIndex])
// 		fmt.Println("animals value #2:", animals[secondIndex])

		 for secondIndex + 1 <= len(animals) {

		 	var firstString string = animals[firstIndex]
      var secondString string = animals[secondIndex]

       if strings.ToLower(firstString) > strings.ToLower(secondString) {
         animals[firstIndex] = secondString
         animals[secondIndex] = firstString
       }

		firstIndex++
		secondIndex++
		}

}

func bubbleSort(animals []string) {
   var N int = len(animals)
   var i int
   for i = 0; i < N; i++ {
     sweep(animals)
   }
 }

func main() {
	var animals []string = []string{
		"dog",
		"cat",
		"alligator",
		"cheetah",
		"rat",
		"moose",
		"cow",
		"mink",
		"porcupine",
		"dung beetle",
		"camel",
		"steer",
		"bat",
		"hamster",
		"horse",
		"colt",
		"bald eagle",
		"frog",
		"rooster",
	}

	fmt.Println("Unsorted:", animals)
	//sweep(animals)
	bubbleSort(animals)
	fmt.Println("Sorted:  ", animals)
}

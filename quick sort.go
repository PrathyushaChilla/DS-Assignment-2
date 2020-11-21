program1

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	quicksort(slice)
	fmt.Println("\n--- Sorted ---\n", slice)
}
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
 
func quicksort(a []int) []int {
    if len(a) < 2 {
        return a
    }
     
    left, right := 0, len(a)-1
     
    pivot := rand.Int() % len(a)
     
    a[pivot], a[right] = a[right], a[pivot]
     
    for i, _ := range a {
        if a[i] < a[right] {
            a[left], a[i] = a[i], a[left]
            left++
        }
    }
     
    a[left], a[right] = a[right], a[left]
     
    quicksort(a[:left])
    quicksort(a[left+1:])
     
    return a
}
output:
--- Unsorted --- 

 [-86 -816 -205 86 -106 -624 357 189 -263 149 419 125 -277 -72 235 117 -361 -321 528 -170]

--- Sorted ---
 [-816 -624 -361 -321 -277 -263 -205 -170 -106 -86 -72 86 117 125 149 189 235 357 419 528]

Program exited.









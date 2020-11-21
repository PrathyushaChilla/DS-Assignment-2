program1

package main

import (
    "fmt"
)

func main() {
    A := []int{3, 5, 1, 6, 1, 7, 2, 4, 5}
    fmt.Println(sort(A))
}
func sort(A []int) []int {
    if len(A) <= 1 {
        return A
    }

    left, right := split(A)
    left = sort(left)
    right = sort(right)
    return merge(left, right)
}

func split(A []int) ([]int, []int) {
    return A[0:len(A) / 2], A[len(A) / 2:]
}
func merge(A, B []int) []int {
    arr := make([]int, len(A) + len(B))
    j, k := 0, 0

    for i := 0; i < len(arr); i++ {
                   arr[i] = B[k]
            k++
            continue
        } else if k >= len(B) {
            arr[i] = A[j]
            j++
            continue
        }
                if A[j] > B[k] {
            arr[i] = B[k]
            k++
        } else {
            arr[i] = A[j]
            j++
        }
    }

    return arr
}
output:
[1 1 2 3 4 5 5 6 7]

Program exited.
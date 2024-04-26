package main

import "fmt"

func BubbleSort(sli []int) {
	for i := range sli {
		for j := 0; j < len(sli)-1-i; j++ {
			if sli[j] > sli[j+1] {
				sli[j], sli[j+1] = sli[j+1], sli[j]
			}
		}
	}
}

func main() {
	sli := []int{32, 3, 6, 132, 9, 0, -1, 43, 57, 5, 3, 2}
	BubbleSort(sli)
	fmt.Println(sli)
}

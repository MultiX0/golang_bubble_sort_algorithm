package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

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
	now := time.Now()
	defer func() {
		fmt.Println("result in :", time.Since(now))
	}()
	sli := []int{}
	var length, input int

	fmt.Println("What is the length of the array that you want ?")
	fmt.Print(" > ")
	fmt.Scan(&input)
	length = input

	for i := 0; i < length; i++ {
		fmt.Println("value", i+1)
		fmt.Print(" > ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			num := scanner.Text()
			parsedNum, _ := strconv.Atoi(num)
			sli = append(sli, parsedNum)
		}
	}

	fmt.Println("Sorted List :")
	BubbleSort(sli)
	fmt.Println(sli)
}

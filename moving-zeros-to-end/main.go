package main

import "fmt"

func main() {
	fmt.Println(MoveZeros([]int{1, 0, 1, 0, 3, 0, 0, 0, 0, 1, 2, 0}))
}

func MoveZeros(arr []int) []int {
	zerosFlag := false
	zeroIndex := 0
	for i := 0; i <= len(arr)-1; i++ {
		if arr[i] == 0 {
			if !zerosFlag {
				zeroIndex = i
				zerosFlag = true
			}
			continue
		}
		zerosFlag = false
		exchangeNum := arr[zeroIndex]
		arr[zeroIndex] = arr[i]
		arr[i] = exchangeNum
		zeroIndex++
		if len(arr)-1 >= i+1 {
			if arr[i+1] == 0 {
				zerosFlag = true
			}
		}
	}
	return arr
}

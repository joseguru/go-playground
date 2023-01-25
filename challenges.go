package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println(sevenBoom([]int{1, 2, 87}))

}

func sevenBoom(myslice []int) string {
	for _, s := range myslice {
		if checkDigit(s) {
			return "Boom!"
		}
	}
	return "there is no 7 in the array"
}
func checkDigit(dig int) bool {
	str := strconv.Itoa(dig)

	if strings.Contains(str, "7") {
		return true
	}

	return false
}

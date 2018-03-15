package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func ReadInput(acc []string) []string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return ReadInput(append(acc, scanner.Text()))
	} else {
		return acc
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input: ", err)
	}
	return acc
}

func ListToNumber(acc int, i int, numbers []string) int {
	if i == len(numbers) {
		return acc
	}
	num, err := strconv.Atoi(numbers[i])
	if err != nil {
		fmt.Fprintln(os.Stderr, "erro converting string to int: ", err)
	}
	if num > 0 {
		acc := acc + num*num
		return ListToNumber(acc, i+1, numbers)
	} else {
		return ListToNumber(acc, i+1, numbers)
	}
	return acc
}

func IterateList(i int, acc []int, numbers []string) []int {
	if i == len(numbers) {
		return acc
	}
	if len(numbers[i]) > 1 {
		listToSum := strings.Split(numbers[i], " ")
		num := ListToNumber(0, 0, listToSum)
		fmt.Println(num)
		return IterateList(i + 1, append(acc, num), numbers)
	}
	return IterateList(i + 1, acc, numbers)
}


func main(){
	var inputAcc []string
	var numsAcc []int 
	numbers := ReadInput(inputAcc)
	nums := IterateList(0, numsAcc, numbers)
	_ = nums
}

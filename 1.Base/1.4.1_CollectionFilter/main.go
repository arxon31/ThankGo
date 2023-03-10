package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func filter(predicate func(int) bool, iterable []int) []int {
	// отфильтруйте `iterable` с помощью `predicate`
	// и верните отфильтрованный срез
	var filteredSlice []int
	var iterableValue int
	for _, iterableValue = range iterable {
		if predicate(iterableValue) == true {
			filteredSlice = append(filteredSlice, iterableValue)
		}
	}
	return filteredSlice
}

func main() {
	src := readInput()
	// отфильтруйте `src` так, чтобы остались только четные числа
	// и запишите результат в `res`
	// res := filter(...)
	res := filter(func(i int) bool {
		if i%2 == 0 {
			return true
		} else {
			return false
		}
	}, src)
	fmt.Println(res)
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// readInput считывает целые числа из `os.Stdin`
// и возвращает в виде среза
// разделителем чисел считается пробел
func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return nums
}

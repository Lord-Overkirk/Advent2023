package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func is_num(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func line_solution(line string) (res int) {
	var pos_x, pos_y rune
	for _, char := range line {
		if is_num(byte(char)) {
			if pos_x == '\x00' {
				pos_x = char
			} else {
				pos_y = char
			}
		}
	}
	if pos_y == '\x00' {
		pos_y = pos_x
	}
	var chars [] rune
	chars = append(chars, pos_x)
	chars = append(chars, pos_y)
	line_sol := string(chars)
	res, err := strconv.Atoi(line_sol)

	if err != nil {
        log.Fatal(err)
	}
	return
}

func sum_of_input(input [] string) (sum int) {
	for _, line := range input {
		sum += line_solution(line)
	}
	return
}

func main() {
	file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }

	scanner := bufio.NewScanner(file)

	var input_list [] string
	for scanner.Scan() {
		input_list = append(input_list, scanner.Text())
	}

	sol := sum_of_input(input_list)
	fmt.Println(sol)
}
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func is_num(char rune) bool {
	if char >= '0' && char <= '9' {
		return true
	}
	return false
}

func is_symbol(char rune) bool {
	if char != '.' && !(char >= '0' && char <= '9') {
		return true
	}
	return false
}

func get_number(i int, j int, input [] string) (res int) {
	var digits [] rune
	var current_digit rune = rune(input[i][j])
	for is_num(current_digit) {
		digits = append(digits, current_digit)

		if (j+1) > len(input[i]) - 1 {
			break
		}
		j++
		current_digit = rune(input[i][j])
	}
	res, err := strconv.Atoi(string(digits))

	if err != nil {
		log.Fatal("Ätoi failed")
	}

	return res
}

func get_num_digits(i int, j int, input [] string) int {
	var start_j = j
	var current_digit rune = rune(input[i][j])
	for is_num(current_digit) {
		if (j+1) > len(input[i]) - 1 {
			break
		}
		j++
		current_digit = rune(input[i][j])
	}
	return j - start_j
}

func is_engine_part(i int, j int, input [] string) bool {
	var num_digits int = get_num_digits(i, j, input)
	start_j := j-1
	end_j := j + num_digits + 1
	start_i := i-1
	end_i := i+1

	if (i-1) < 0 {
		start_i = i
	}
	if (i+1) >= len(input) {
		end_i = len(input) - 1
	}
	if start_j < 0 {
		start_j++
	}
	if end_j > len(input[i]) {
		end_j--
	}
	
	for cur_i := start_i; cur_i <= end_i; cur_i++ {
		for cur_j := start_j; cur_j < end_j; cur_j++ {
			to_eval := rune(input[cur_i][cur_j])
			if is_symbol(to_eval) {
				return true
			}
		}
	}
	return false
}

func get_part_value(i int, j int, input [] string) (res int) {
	if (j-1) >= 0 {
		var prev_digit rune = rune(input[i][j-1])
		// Already processed 
		if is_num(prev_digit) {
			return 0
		}
	}

	if is_engine_part(i, j, input) {
		res += get_number(i, j, input)
	}
	
	return
}

func solution_3a(input [] string) (res int) {
	for i, line := range input {
		for j, char := range line {
			if is_num(char) {
				res += get_part_value(i, j, input)
			}
		}
	}
	return
}

func has_gear_ratio(i int, j int, input [] string) bool {
	candidate := 0

	start_i := i-1
	end_i := i+1
	start_j := j-1
	end_j := j+1

	if (i-1) < 0 {
		start_i = i
	}
	if (i+1) >= len(input) {
		end_i = len(input) - 1
	}

	if end_j > len(input[i]) - 1 {
		end_j--
	}

	for cur_i := start_i; cur_i <= end_i; cur_i++ {
		for cur_j := start_j; cur_j <= end_j; cur_j++ {
			if is_num(rune(input[cur_i][cur_j])) {
				candidate++
				for cur_j+1 < len(input[cur_i]) {
					if !is_num(rune(input[cur_i][cur_j])) {
						break
					}
					cur_j++
				}
			}
		}
	}
	if candidate == 2 {
		return true
	}
	return false
}

func get_first_digit(input string, j int) int {
	current := rune(input[j])
	for is_num(current) {
		if (j-1) < 0 {
			return j
		}
		j--
		current = rune(input[j])
	}
	return j+1
}

func gear_ratio(i int, j int, input [] string) (res int) {
	if !has_gear_ratio(i, j, input) {
		return 0
	}

	start_i := i-1
	end_i := i+1
	start_j := j-1
	end_j := j+1

	if (i-1) < 0 {
		start_i = i
	}
	if (i+1) >= len(input) {
		end_i = len(input) - 1
	}

	if end_j > len(input[i]) - 1 {
		end_j--
	}

	var gears [] int
	for cur_i := start_i; cur_i <= end_i; cur_i++ {
		for cur_j := start_j; cur_j <= end_j; cur_j++ {
			if is_num(rune(input[cur_i][cur_j])) {
				first_j := get_first_digit(input[cur_i], cur_j)
				num := get_number(cur_i, first_j, input)
				gears = append(gears, num)
				for cur_j+1 < len(input[cur_i]) {
					if !is_num(rune(input[cur_i][cur_j])) {
						break
					}
					cur_j++
				}
			}
		}
	}

	if len(gears) != 2 {
		fmt.Println(gears)
		log.Fatal("Incorrect gears")
	}
	res = gears[0] * gears[1]
	return
}

func solution_3b(input [] string) (res int) {
	for i, line := range input {
		for j, char := range line {
			if char == '*' {
				res += gear_ratio(i, j, input)
			}
		}
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

	res_3a := solution_3a(input_list)
	fmt.Println(res_3a)

	res_3b := solution_3b(input_list)
	fmt.Println(res_3b)
}
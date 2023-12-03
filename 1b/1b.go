package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var number_map = map[string]int{
    "one"	: 1,
    "two"	: 2,
    "three"	: 3,
    "four"	: 4,
    "five"	: 5,
    "six"	: 6,
    "seven"	: 7,
    "eight"	: 8,
    "nine"	: 9,
}

type StringSol struct {
	number int
	idx int
}

type IntSol struct {
	number int
	idx int
}

func is_num(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func index_per_substring(input string, target string) (indexes [] int) {
	for i, char := range input {
		if char == rune(target[0]) {
			for j, rest := range target {
				if i + j < len(input) {
					if rune(input[i+j]) != rest {
						break;
					}
					if j == len(target) -1 {
						indexes = append(indexes, i)
					}
				}
			}
		}
	}
	return
}

func substring_check(input_line string) (solutions [] StringSol) {
	for key := range number_map {
		idxs := index_per_substring(input_line, key)
		for _, idx := range idxs {
			solutions = append(solutions, StringSol{number_map[key], idx})
		}
	}
	sort.Slice(solutions, func (i, j int) bool {
		return solutions[i].idx < solutions[j].idx
	})
	return
}

func digit_check(input_line string) (solutions [] IntSol) {
	for idx, char := range input_line {
		if is_num(char) {
			digit:= int(char) - 48
			curr_sol := IntSol{digit, idx}
			solutions = append(solutions, curr_sol)
		}
	}
	return
}

func sum_of_input(input [] string) (sum int) {
	var string_sols [] StringSol
	var int_sols [] IntSol

	for _, line := range input {
		// fmt.Printf(line)
		string_sols = substring_check(line)
		int_sols = digit_check(line)

		var first_string, last_string StringSol
		var first_digit, last_digit IntSol
		if len(string_sols) != 0 {
			first_string = string_sols[0]
			last_string = string_sols[len(string_sols)-1]
		}
		if len(int_sols) != 0 {
			first_digit = int_sols[0]
			last_digit = int_sols[len(int_sols)-1]
		}
		if len(string_sols) == 0 {
			first_string = StringSol(first_digit)
			last_string = StringSol(last_digit)
		}
		if len(int_sols) == 0 {
			first_digit = IntSol(first_string)
			last_digit = IntSol(last_string)
		}
		
		var pos_x, pos_y string

		if first_string.idx < first_digit.idx {
			pos_x = strconv.Itoa(first_string.number)
		} else {
			pos_x = strconv.Itoa(first_digit.number)
		}

		if last_string.idx < last_digit.idx {
			pos_y = strconv.Itoa(last_digit.number)
		} else {
			pos_y = strconv.Itoa(last_string.number)
		}
		if pos_x == "0" {
			pos_x = pos_y
		} else if pos_y == "0" {
			pos_y = pos_x
		}

		line_sol := pos_x + pos_y
		res, err := strconv.Atoi(line_sol)
		if err != nil {
			log.Fatal(err)
		}
		sum += res
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
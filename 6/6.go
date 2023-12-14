package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parse_inputs(input string) (res [] int) {
	prefix_del := strings.Split(input, ":")
	no_whitespace := strings.TrimSpace(prefix_del[1])

	input_nums := strings.Split(no_whitespace, " ")
	big_num := strings.Join(input_nums, "")

	for _, num := range input_nums {
		conv_num, _ := strconv.Atoi((num))

		if conv_num != 0 {
			res = append(res, conv_num)
		}
	}

	conv_num, _ := strconv.Atoi((big_num))
	res = []int{conv_num}

	return
}

func get_num_winning_races(time int, distance int) (wins int) {
	for i := 0; i < time; i++ {
		velo := i
		remaining_time := time - i

		achieved_dist := velo*remaining_time

		if achieved_dist > distance {
			wins++
		}
	}
	return
}

func sol_6a(times [] int, distances [] int) (res int) {
	var win_list [] int
	for i := 0; i < len(times); i++ {
		time := times[i]
		distance := distances[i]

		win_list = append(win_list, get_num_winning_races(time, distance)) 
	}

	res = 1
	for _, win := range win_list {
		res *= win
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

	times := input_list[0]
	distances := input_list[1]
	
	parsed_times := parse_inputs(times)
	parsed_distances := parse_inputs(distances)

	fmt.Println(parsed_times)
	fmt.Println(parsed_distances)

	prod := sol_6a(parsed_times, parsed_distances)
	fmt.Println(prod)
}
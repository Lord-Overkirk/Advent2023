package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	winning [] int
	numbers [] int
}

func parse_card(raw_card string) (res Card) {
	game := strings.Split(raw_card, ":")
	raw_game := game[1]
	parsed_card := strings.Split(raw_game, "|")

	raw_winning := strings.TrimSpace(parsed_card[0])
	raw_numbers := strings.TrimSpace(parsed_card[1])
	
	// for _ 
	raw_winning_list := strings.Split(raw_winning, " ")
	raw_numbers_list := strings.Split(raw_numbers, " ")

	for _, winning_num := range raw_winning_list {
		if winning_num == "" {
			continue
		}

		parsed_num, err := strconv.Atoi(winning_num)
		
		if err != nil {
			log.Fatal("Atoi failed")
		}

		res.winning = append(res.winning, parsed_num)
	}
	
	for _, card_number := range raw_numbers_list {
		if card_number == "" {
			continue
		}

		parsed_num, err := strconv.Atoi(card_number)

		if err != nil {
			log.Fatal("Atoi failed")
		}

		res.numbers = append(res.numbers, parsed_num)
	}

	return
}

func get_points_per_card(card Card) (points int) {
	for _, num := range card.numbers {
		for _, win_num := range card.winning {
			if num == win_num {
				if points == 0 {
					points++
				} else {
					points*=2
				}
			}
		}
	}
	return
}

func solution_4a(input [] string) (res int) {
	for _, card := range input {
		parsed_card := parse_card(card)
		res += get_points_per_card(parsed_card)
	}
	return
}

func get_num_matches(card Card) (matches int) {
	for _, num := range card.numbers {
		for _, win_num := range card.winning {
			if num == win_num {
				matches++
			}
		}
	}
	return
}

func sum_slice(slice [] int) (sum int) {
	for _, elem := range slice {
		sum += elem
	}
	return
}

func solution_4b(input [] string) (res int) {
	match_list := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		match_list[i] = 1
	}
	for card_num, card := range input {
		parsed_card := parse_card(card)
		matches := get_num_matches(parsed_card)
		// for copies := 1; copies < match_list[card_num]; copies++ {
		// 	// matches *= 2
		// }
		for match := 1; match <= matches; match++ {
			// fmt.Println("matches for card", card_num, matches, match_list[card_num])
			if card_num+match < len(match_list) {
				for copies := 1; copies <= match_list[card_num]; copies++ {
					match_list[card_num+match]++
				}
			}
		}
	}
	res = sum_slice(match_list)
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

	res_4a := solution_4a(input_list)
	fmt.Println(res_4a)

	res_4b := solution_4b(input_list)
	fmt.Println(res_4b)
}
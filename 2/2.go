package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	game_id 	int
	rounds[]	Round
}

type Round struct {
	num_reds 	int
	num_greens 	int
	num_blues 	int
}

const MAX_RED int = 12
const MAX_GREEN int = 13
const MAX_BLUE int = 14


func parse_color(colors [] string) (round Round) {
	for _, color := range colors {
		trimmed := strings.TrimSpace(color)
		parsed := strings.Split(trimmed, " ")
		n := parsed[0]
		c := parsed[1]
		cube_amount, err := strconv.Atoi(n)

		if err != nil {
			log.Fatal("Error in atoi")
		}

		if c[0] == 'r' {
			round.num_reds = cube_amount
		} else if c[0] == 'g' {
			round.num_greens = cube_amount
		} else if c[0] == 'b' {
			round.num_blues = cube_amount
		} else {
			log.Fatal("Error parsing color")
		}
	}

	return
}

func parse_games(games [] string) (parsed_games [] Game) {
	for i, game := range games {
		var cur_game Game
		
		parts := strings.Split(game, ":")
		cur_game.game_id = i+1
		
		rounds := strings.Split(parts[1], ";")
		for _, round := range rounds {
			color := strings.Split(round, ",")
			parsed_round := parse_color(color)
			cur_game.rounds = append(cur_game.rounds, parsed_round)
		}
		parsed_games = append(parsed_games, cur_game)
	}
	return
}

func game_is_possible_part_a(cur_game Game) bool {
	for _, round := range cur_game.rounds {
		if round.num_reds > MAX_RED || round.num_blues > MAX_BLUE || round.num_greens > MAX_GREEN {
			return false
		}
	}
	return true
}

func game_solution(games [] Game) (res int) {
	for _, game := range games {
		if game_is_possible_part_a(game) {
			res += game.game_id
		}
	}
	return
}

func game_solution_2b(games [] Game) (res int) {
	for _, game := range games {
		var min_red, min_green, min_blue int
		for _, round := range game.rounds {
			if round.num_reds > min_red {
				min_red = round.num_reds
			}
			if round.num_blues > min_blue {
				min_blue = round.num_blues
			}
			if round.num_greens > min_green {
				min_green = round.num_greens
			}
		}
		pow := min_red * min_green * min_blue
		res += pow
	}
	return
}

func main()  {
	file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }

	scanner := bufio.NewScanner(file)

	var input_list [] string
	for scanner.Scan() {
		input_list = append(input_list, scanner.Text())
	}
	games := parse_games(input_list)
	// result := game_solution(games)
	result2 := game_solution_2b(games)
	fmt.Println(result2)
}
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type FarmRange struct {
	dest_start uint64
	source_start uint64

	dest_end uint64
	source_end uint64
}

func init_farm_range(raw_map [] uint64) (fr FarmRange) {
	if len(raw_map) != 3 {
		log.Fatal("Wrong farm range")
	}
	dest := raw_map[0]
	src := raw_map[1]
	len := raw_map[2]

	// var dests, srcs [] int
	// for i := 0; i < len; i++ {
	// 	dests = append(dests, dest+i)
	// 	srcs = append(srcs, src+i)
	// }
	fr.dest_start = dest
	fr.dest_end = dest + len
	fr.source_start = src
	fr.source_end = src + len

	return
}

func (f FarmRange) get(i uint64) uint64 {
	if i >= f.source_start && i <= f.source_end {
		difference := i - f.source_start
		// fmt.Printf("%d -> %d\n", i, f.dest_start + difference)
		return f.dest_start + difference
	} else {
		// fmt.Printf("%d -> %d\n", i, i)
		return i
	}
	// for j, target := range f.source{
	// 	if target == i {
	// 		return f.dest[j]
	// 	}
	// }
	// return i
}

func get_seeds(input string) (seeds [] uint64) {
	trimmed := strings.Replace(input, "seeds: ", "", 1)

	for _, char := range strings.Split(trimmed, " ") {
		seed, err := strconv.ParseUint(char, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		seeds = append(seeds, seed)
	}
	return
}

func parse_ranges(input string) (ranges [] uint64) {
	for _, char := range strings.Split(input, " ") {
		num, err := strconv.ParseUint(char, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		ranges = append(ranges, num)
	}
	return
}

func parse_mapping(input [] string) [][][]uint64 {
	mappings := make([][][]uint64, 7)
	var num_mappings = 0

	for _, line := range input {
		if len(line) == 0 {
			num_mappings++
			continue
		}
		if line != "" && (line[0] >= '0' && line[0] <= '9') {
			current_mapping := parse_ranges(line)
			mappings[num_mappings] = append(mappings[num_mappings], current_mapping)
		}
	}
	return mappings
}

func solution_5a(seeds[] uint64, mappings [][][] uint64) (minimum uint64) {
	var locations [] uint64
	for _, seed := range seeds {
		for _, mapping := range mappings {
			for _, current_range := range mapping {
				// fmt.Println(seed)
				var farm_range FarmRange
				if seed >= current_range[1] && seed <= current_range[1] + current_range[2] {
					farm_range = init_farm_range(current_range)
				}
				old_seed := seed
				seed = farm_range.get(seed)
				if old_seed != seed {
					break
				}
			}
		}
		locations = append(locations, seed)
	}
	minimum = slices.Min(locations)
	return 
}

func solution_5b(seed uint64, mappings [][][] uint64) (minimum uint64) {
	for _, mapping := range mappings {
		// fmt.Printf("%d %d -> ", i, seed)
		for _, current_range := range mapping {
			var farm_range FarmRange
			if seed >= current_range[1] && seed <= current_range[1] + current_range[2] {
				farm_range = init_farm_range(current_range)
			}
			old_seed := seed
			seed = farm_range.get(seed)
			// fmt.Println(farm_range, seed)
			if old_seed != seed {
				break
			}
		}
	}
	minimum = seed
	return
}

func spawn_5b(seed_start uint64, seed_end uint64, mappings [][][] uint64, c chan uint64) {
	var min uint64 = 9999999999999
	for j := seed_start; j < seed_end; j++ {
		new_min := solution_5b(j, mappings)
		if new_min < min {
			min = new_min
		}
	}
	c <- min
}

func calc_seed(mappings [][][] int) (smallest_map [] int) {
	sm := make([][]int, 7)
	smallest_dest := 9999999999999
	for _, mapping := range mappings {
		for _, current_range := range mapping {
			dst := current_range[0]
			// src := current_range[1]
			// len := current_range[2]
			if dst <= smallest_dest {
				smallest_dest = dst
				// smallest_source = src
				smallest_map = current_range
				sm = append(sm, smallest_map)
			}
		}
	}
	// fmt.Println(sm)
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
	seeds_5a := get_seeds(input_list[0])
	
	
	maps := parse_mapping(input_list[2:])
	min := solution_5a(seeds_5a, maps)
	fmt.Println(min)

	fmt.Println("--- part b ---")
	min = 9999999999999
	c := make(chan uint64)
	var minima [] uint64
	for i := 0; i < len(seeds_5a)-1; i+=2 {
		// fmt.Println(i)
		go spawn_5b(seeds_5a[i], (seeds_5a[i] + seeds_5a[i+1]), maps, c)
	}
	for i := 0; i < len(seeds_5a)/2; i++ {
		// fmt.Println(i, minima)
		minima = append(minima, <-c)
	}
	fmt.Println(minima, slices.Min(minima))
}
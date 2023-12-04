package main

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
}
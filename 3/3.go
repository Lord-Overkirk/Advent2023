package main

func is_symbol(char string) bool {
	if char != "." || (char < "0"&& char > "9") {
		return true
	}
	return false
}

func solution_3a() (res int) {

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
}
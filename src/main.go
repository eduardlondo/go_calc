package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func getGOMAXPROCS() int {
	return runtime.GOMAXPROCS(0)
}

func main() {
	fmt.Printf("Core availables are %d\n", getGOMAXPROCS())
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		line, _ := reader.ReadString('\n')
		aux := removeCharsByASCII(line, 13, 10, 0)
		if aux == "exit" {
			break
		}
		result := eval(line)
		fmt.Println(result)
	}
}

func removeCharsByASCII(str string, chars ...int) string {
	// Convert the string to a rune slice
	runes := []rune(str)

	// Remove the specified characters
	j := 0
	for _, r := range runes {
		if !contains(chars, int(r)) {
			runes[j] = r
			j++
		}
	}
	runes = runes[:j]

	// Convert the rune slice back to a string
	return string(runes)
}

// Helper function to check if a slice contains a value
func contains(s []int, val int) bool {
	for _, v := range s {
		if v == val {
			return true
		}
	}
	return false
}

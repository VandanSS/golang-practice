package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a string: ")
	str, _ := in.ReadString('\n')
	
	str = strings.ToLower(string(str))
	str = strings.TrimSpace(str)
	strlen := len(str)
	
	if string(str[0]) == "i" && string(str[strlen - 1]) == "n" && strings.Contains(str, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
	
}
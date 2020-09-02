package main

import (
	"encoding/json"
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	personMap := make(map[string]string)

	fmt.Println("Enter name : ")
	name, err := in.ReadString('\n')
	name = strings.TrimRight(name, "\r\n")

	fmt.Println("Enter Address : ")
	addr, err := in.ReadString('\n')
	addr = strings.TrimRight(addr, "\r\n")

	personMap["name"] = name
	personMap["address"] = addr

	personJson, err := json.Marshal(personMap)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(personJson))
	}
}
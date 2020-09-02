package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
    maxLength = 20
)

type name struct {
	fname string
	lname string
}

func main() {
	var fd string
	names := make([]name, 0)
	
	fmt.Println("Enter File Name with extension : ")
	_, err := fmt.Scan(&fd)
	
	if err == nil {
	
		fileData, _ := ioutil.ReadFile(fd)
		
		allNames := strings.Split(string(fileData), "\n")
		
		for _, person := range allNames {
			temp := strings.Split(person, " ")
			
			if len(temp[0]) > maxLength {
				rs := []rune(temp[0])
				temp[0] = string(rs[:maxLength])
			}
			
			if len(temp[1]) > maxLength {
				rs := []rune(temp[1])
				temp[1] = string(rs[:maxLength])
			}
			
			newName := name{temp[0], temp[1]}
			names = append(names, newName)
		}
		
		for _, nm := range names {
			fmt.Println(nm.fname, nm.lname)
		}
		
	} else {
		fmt.Println(err)
	}
}
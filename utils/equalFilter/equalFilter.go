package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("dump.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	//go through each line in a file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var url string = scanner.Text()
		if strings.Contains(url, "=") {
			fmt.Println(url)
		}
	}

}

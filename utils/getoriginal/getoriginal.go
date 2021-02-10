package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var url string
var url2 string

func scan2() {
	file, err := os.Open("haveParam.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner2 := bufio.NewScanner(file)
	for scanner2.Scan() {
		url2 = scanner2.Text()
		if strings.Contains(url2, url) {
			fmt.Println(url2)
			return

		}
	}

}

func main() {

	newfile, err := os.Open("new.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer newfile.Close()

	//go through each line in a file
	scanner := bufio.NewScanner(newfile)

	for scanner.Scan() {
		url = scanner.Text()
		url = (strings.Split(url, "="))[0]
		scan2()
	}

}

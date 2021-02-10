package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var param []string

func main() {
	file, err := os.Open("haveParam.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	//go through each line in a file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var url string = scanner.Text()

		//fmt.Println(url)
		//fmt.Println("New URL discovered! : ", url)
		splitURL := strings.Split(url, "=")
		//fmt.Println("SPLIT URL : ", splitURL)
		if len(splitURL) > 2 {
			for i := 0; i < len(splitURL); i++ {
				if strings.Contains(splitURL[i], "&") {
					temp := strings.Split(splitURL[i], "&")
					param = append(param, temp[0])
				} else {
					param = append(param, splitURL[len(splitURL)-1])
				}
			}
		} else {
			param = append(param, splitURL[len(splitURL)-1])
		}

		//get clean url for db
		cleanURL := url
		for i := 0; i < len(param); i++ {
			cleanURL = strings.Replace(cleanURL, param[i], "", -1)
		}
		//fmt.Println("PARAM ARRAY : ", param)
		//fmt.Println("URL : ", url)
		fmt.Println(cleanURL)
		param = nil
	}
}

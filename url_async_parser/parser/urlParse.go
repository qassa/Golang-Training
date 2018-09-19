package parser

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func ParseFile(nameFile string) ([]string, error) {
	file, err := os.Open(nameFile)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scannerUrls := bufio.NewScanner(file)
	var urls []string
	for scannerUrls.Scan() {
		urls = append(urls, scannerUrls.Text())
	}
	return urls, nil
}

func GetAndWrite(url string, index int, c chan func() (int, int)) {
	resp, err := http.Get(url)
	defer (func() {
		if resp != nil {
			resp.Body.Close()
		}
	})()

	if err != nil || resp == nil {
		fmt.Println("Error while requesting the page %s: %s", url, err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	c <- (func() (int, int) { return index, len(body) })
	if err != nil {
		fmt.Errorf("Error of reading response: %s\n", err.Error())
	}
	file, err := os.Create(strconv.Itoa(index) + ".html")

	if err != nil {
		fmt.Printf("Error while creating the file %s\n", err.Error())
	}
	file.Write(body)
	file.Close()
}

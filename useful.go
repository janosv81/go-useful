package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
log.Fatal("Check the source code ;)")
}

func readLinesFromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func batchPostWithBearer(urls []string, bearerToken string) error {
	for index, url := range urls {
		fmt.Print(index)
		fmt.Print(". - ")
		fmt.Println(url)
		req, _ := http.NewRequest("POST", url, nil)

		req.Header.Add("Authorization", "Bearer "+bearerToken)

		res, err := http.DefaultClient.Do(req)
		if res.StatusCode != 200 {
			return fmt.Errorf("Response NOT OK! URL:%s StatusCode:%d", url, res.StatusCode)
		}
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(res.Status)
		fmt.Println(string(body))
	}
	return nil
}

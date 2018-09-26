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

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

//MapReduce reduce function for arrays
func MapReduce(list []ExampleStruct, f func(ExampleStruct) string) []string {
	vsm := make([]string, len(list))
	for i, v := range list {
		vsm[i] = f(v)
	}
	return removeDuplicates(vsm)
}

func removeDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

//Filter filters arrays by function
func Filter(vs []ExampleStruct, f func(ExampleStruct) bool) []ExampleStruct {
	vsf := make([]ExampleStruct, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vs, v);
		}
	}
	return vsf
}

func keyOfMaxValue(input map[string]int) string {
	max := 0
	maxKey := ""
	for k, v := range input {
		if v > max {
			max = v
			maxKey = k
		}
	}
	return maxKey
}

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Urls struct {
	Page string
}

func main() {
	readFile := os.Args[1]
	writeFile := os.Args[2]

	data, err := ioutil.ReadFile(readFile)
	if err != nil {
		log.Fatal(err)
	}

	urls := parseIncomingCsv(string(data))

	file, err := os.Create(writeFile)
	if err != nil {
		log.Fatalln("Error creating file", err)
	}

	defer file.Close()

	l := dataToWriteToCsv(urls)

	w := csv.NewWriter(file)
	defer w.Flush()
	w.WriteAll(l)

	if err := w.Error(); err != nil {
		log.Fatalln("Error writing to csv", err)
	}

	fmt.Println("Process finished successfully")
}

func dataToWriteToCsv(urls []Urls) [][]string {
	m := make(map[string]interface{})
	var l [][]string
	for x := 0; x < len(urls); x++ {
		if _, ok := m[urls[x].Page]; ok {

		} else {
			m[urls[x].Page] = nil
			u := []string{
				urls[x].Page,
			}
			l = append(l, u)
		}

	}
	return l
}

func parseIncomingCsv(data string) []Urls {
	var urls []Urls

	r := csv.NewReader(strings.NewReader(data))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Error reading csv", err)
		}
		var url = Urls{
			record[0],
		}

		urls = append(urls, url)

	}
	return urls
}

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func readCsv(csvfile string) {
	file, err := os.Open(csvfile)
	if err != nil {
		log.Fatalf("Failed opening file")
		panic(err)
	}

	defer file.Close()
	csvReader := csv.NewReader(file)

	for {
		csvRecord, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println("CSV Record")
		for _, csvValue := range csvRecord {
			fmt.Println(csvValue)
		}
	}
	/*scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var csvlines []string

	for scanner.Scan() {
		csvlines = append(csvlines, scanner.Text())
	}

	for _, eachline := range csvlines {
		fmt.Println(eachline)
	}*/

}

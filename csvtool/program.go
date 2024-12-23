package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fileName := os.Args[1]
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatal("Unable to open file:", err)
		}
		defer file.Close()

		reader := csv.NewReader(file)

		records, err := reader.ReadAll()

		if err != nil {
			log.Fatal("Unable to read file:", err)
		}

		for _, row := range records {
			fmt.Println(row)
		}
	}
}

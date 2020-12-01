package sumupper

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func SumUp(path_data string) {
	files, err := ioutil.ReadDir(path_data)

    if err != nil {
        log.Fatal(err)
	}
	
    for _, file := range files {
		if strings.Contains(file.Name(), "csv") {
			get_total_rainfall(file.Name())
		}
    }
}

func get_total_rainfall(path_file string) {
	RAINFALL_COLUMN := 2
	csvfile, err := os.Open(path_file)

	if err != nil {
		log.Fatalln("Unable to open CSV file", err)
	}

	reader := csv.NewReader(csvfile)

	rainfall_total := 0.000
	num_values := 0

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Failed to parse CSV file", err)
		}
		
		if rainfall, err := strconv.ParseFloat(record[RAINFALL_COLUMN], 64); err == nil {
			rainfall_total += rainfall
		}
		num_values += 1
	}

	fmt.Println("Total values found in CSV: " + strconv.Itoa(num_values) + " Total Rainfall: " + strconv.FormatFloat(rainfall_total, 'f', -1, 64))
}
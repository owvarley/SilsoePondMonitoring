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
	"container/list"
)

type RainfallData struct {
	Total float64
	Date string
}

func SumUp(path_data string) {
	rainfall_data := list.New()
	files, err := ioutil.ReadDir(path_data)

    if err != nil {
        log.Fatal(err)
	}
	
	// Read the values out of the CSV and generate a total and date
    for _, file := range files {
		if strings.Contains(file.Name(), "csv") {
			rainfall_data.PushBack(get_total_rainfall(file.Name()))
		}
	}

	// Write the values to a CSV file
	write_to_csv(rainfall_data)
}
	
func get_total_rainfall(path_file string) RainfallData {
	RAINFALL_DATE := 0
	RAINFALL_COLUMN := 2

	// Open the CSV file for reading
	csvfile, err := os.Open(path_file)

	if err != nil {
		log.Fatalln("Unable to open CSV file", err)
	}

	reader := csv.NewReader(csvfile)

	rainfall_total := 0.000
	date_of_recording := ""
	num_values := 0

	// Parse the contents of the CSV
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Failed to parse CSV file", err)
		}

		date_of_recording = record[RAINFALL_DATE]
		
		if rainfall, err := strconv.ParseFloat(record[RAINFALL_COLUMN], 64); err == nil {
			rainfall_total += rainfall
		}
		num_values += 1		
	}

	fmt.Println("Total values found in CSV: " + strconv.Itoa(num_values) + " Total Rainfall: " + strconv.FormatFloat(rainfall_total, 'f', -1, 64))

	return RainfallData { rainfall_total, date_of_recording }
}

func write_to_csv(rainfall_data *list.List) {
	for r := rainfall_data.Front(); r != nil; r = r.Next() {
		rainfall_data := r.Value.(RainfallData)
		fmt.Println(rainfall_data.Date + "," + strconv.FormatFloat(rainfall_data.Total, 'f', -1, 64))
	}
}
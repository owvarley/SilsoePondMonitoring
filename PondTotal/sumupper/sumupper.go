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

func (r RainfallData) ToSlice() []string {
	return []string {
		r.Date,
		strconv.FormatFloat(r.Total, 'f', -1, 64)}
}

func SumUp(path_data string) {
	rainfall_data := list.New()
	files, err := ioutil.ReadDir(path_data)

	log.SetOutput(os.Stdout)

	checkError("Failed to read the directory", err)
	
	num_csv_files := 0

	// Read the values out of the CSV and generate a total and date
    for _, file := range files {
		if strings.HasPrefix(file.Name(), "rainfall_data") {
			rainfall_data.PushBack(get_total_rainfall(path_data + file.Name()))
			num_csv_files += 1
		}
	}

	if num_csv_files > 0 {
		// Write the values to a CSV file
		write_to_csv(path_data, rainfall_data)
	} else	{
		fmt.Println("No CSV files found in " + path_data)
	}

}

func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}
	
func get_total_rainfall(path_file string) RainfallData {
	RAINFALL_DATE := 0
	RAINFALL_COLUMN := 2

	// Open the CSV file for reading
	csvfile, err := os.Open(path_file)
	checkError("Unable to open CSV file: ", err)

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
		checkError("Failed to parse CSV file", err)

		date_of_recording = record[RAINFALL_DATE]
		
		if rainfall, err := strconv.ParseFloat(record[RAINFALL_COLUMN], 64); err == nil {
			rainfall_total += rainfall
		}
		num_values += 1		
	}

	fmt.Println("Total values found in CSV: " + strconv.Itoa(num_values) + " Total Rainfall: " + strconv.FormatFloat(rainfall_total, 'f', -1, 64))

	return RainfallData { rainfall_total, date_of_recording }
}

func write_to_csv(path_data string, rainfall_data *list.List) {
	file, err := os.Create(path_data + "total_rainfall.csv")
	checkError("Unable to create CSV file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for r := rainfall_data.Front(); r != nil; r = r.Next() {
		err := writer.Write( r.Value.(RainfallData).ToSlice())
		checkError("Unable to write record into CSV file", err)
	}
}
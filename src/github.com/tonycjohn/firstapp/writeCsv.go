package main

import (
	"encoding/csv"
	"os"
)

func writeCsv(rowCount int, datasetID string, fileName string) {
	//newCsvFile, err:=os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	newCsvFile, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer newCsvFile.Close()

	csvWriter := csv.NewWriter(newCsvFile)
	csvHeader := []string{"version_dataset_id", "dimension_code", "char_code", "char_value_code", "char_value_display_name", "char_value_display_sequence", "insert_time", "update_time"}
	err = csvWriter.Write(csvHeader)
	if err != nil {
		panic(err)
	}

	for i := 0; i <= rowCount; i++ {
		csvRow := []string{datasetID, "dummy_dim", "chummy_char", "villian_value", "nasty_name", "100", "", ""}
		err = csvWriter.Write(csvRow)
		if err != nil {
			panic(err)
		}
		csvWriter.Flush()
	}

}

package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

const (
	viewCsvFileName   = ""
	tableCsvFileName  = "tables_CSNLSSYNDPI_PROD_clean.csv"
	datasetID         = "CSNLSSYNDPI_PROD_3"
	tableSourcePrefix = "gs://bq-ingestion-bucket/CSNLSSYNDPI_PROD_3/"
	location          = "us"
)

//Table represents each BQ tqble
type Table struct {
	TableID             string
	TableSource         string
	ClusterKeys         []string
	PartitionKey        string
	PartitionRangeStart int64
	PartitionRangeEnd   int64
	PartitionInterval   int64
}

//View represents each BQ tqble
type View struct {
	ViewID string
	Query  string
}

//NewDatasetConfig represents the dataset and tables in BQ
type NewDatasetConfig struct {
	DatasetID string
	Location  string
	Tables    []Table
	Views     []View
}

func stringReplace(tableOption bool, viewOption bool) {

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var newConfig NewDatasetConfig
	var newTables []Table
	var newViews []View
	newConfig.DatasetID = datasetID

	//Tables
	if tableOption {
		fmt.Println("Creating Table Config")

		tableCsvFullPath := wd + "\\" + tableCsvFileName
		tableCsvFile, err := os.Open(tableCsvFullPath)
		if err != nil {
			panic(err)
		}
		defer tableCsvFile.Close()

		newConfig.DatasetID = datasetID
		newConfig.Location = location
		tableCsvReader := csv.NewReader(tableCsvFile)
		for {
			csvRecord, err := tableCsvReader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}

			if csvRecord[1] == "name" {
				continue
			}

			tableID := csvRecord[1]
			//storageFolder := strings.ToLower(tableID)
			storageFolder := tableID
			tableSource := tableSourcePrefix + storageFolder + "/*.parquet"
			clusterDefinition := csvRecord[6]

			if len(clusterDefinition) > 0 {
				step1 := strings.ReplaceAll(clusterDefinition, "LINEAR(", "")
				step2 := strings.ReplaceAll(step1, ")", "")
				step3 := strings.TrimSpace(step2)
				clusterKeys := strings.Split(step3, ",")
				newTable := Table{
					TableID:     tableID,
					TableSource: tableSource,
					ClusterKeys: clusterKeys,
				}
				newTables = append(newTables, newTable)
			} else {
				newTable := Table{
					TableID:     tableID,
					TableSource: tableSource,
				}
				newTables = append(newTables, newTable)
			}

		}
		newConfig.Tables = newTables
	}

	//Views
	if viewOption {

		fmt.Println("Creating View Config")

		csvFileFullPath := wd + "\\" + viewCsvFileName
		csvFile, err := os.Open(csvFileFullPath)
		if err != nil {
			panic(err)
		}
		defer csvFile.Close()
		csvReader := csv.NewReader(csvFile)
		for {
			csvRecord, err := csvReader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			if csvRecord[1] == "name" {
				continue
			}

			viewID := csvRecord[1]
			input := csvRecord[7]
			var step2 string
			var step1 string
			if strings.Contains(input, "FROM mt_final_fact_slices") {
				step1 = strings.ReplaceAll(input, "FROM mt_final_fact_slices", "FROM nlsn-connect-data-eng-poc."+datasetID+".MT_FINAL_FACT_SLICES")
				step2 = strings.ReplaceAll(step1, "mt_final_fact_slices.", "")
				step3 := strings.ReplaceAll(step2, "(", "")
				step4 := strings.ReplaceAll(step3, ")", "")
				final := strings.SplitAfterN(step4, "AS", 2)
				viewDDL := final[1]

				newView := View{
					ViewID: viewID,
					Query:  viewDDL,
				}
				newViews = append(newViews, newView)

			} else if strings.Contains(input, "FROM rms_final_fact_slices") {
				step1 = strings.ReplaceAll(input, "FROM rms_final_fact_slices", "FROM nlsn-connect-data-eng-poc."+datasetID+".RMS_FINAL_FACT_SLICES")
				step2 = strings.ReplaceAll(step1, "rms_final_fact_slices.", "")
				step3 := strings.ReplaceAll(step2, "(", "")
				step4 := strings.ReplaceAll(step3, ")", "")
				final := strings.SplitAfterN(step4, "AS", 2)
				viewDDL := final[1]

				newView := View{
					ViewID: viewID,
					Query:  viewDDL,
				}
				newViews = append(newViews, newView)

			} else if strings.Contains(input, "FROM cps_final_fact_slices") {
				step1 = strings.ReplaceAll(input, "FROM cps_final_fact_slices", "FROM nlsn-connect-data-eng-poc."+datasetID+".CPS_FINAL_FACT_SLICES")
				step2 = strings.ReplaceAll(step1, "cps_final_fact_slices.", "")
				step3 := strings.ReplaceAll(step2, "(", "")
				step4 := strings.ReplaceAll(step3, ")", "")
				final := strings.SplitAfterN(step4, "AS", 2)
				viewDDL := final[1]

				newView := View{
					ViewID: viewID,
					Query:  viewDDL,
				}
				newViews = append(newViews, newView)

			}
			newConfig.Views = newViews
		}
	}

	//Write JSON File
	fmt.Println("Writing JSON file")

	file, err := json.MarshalIndent(newConfig, "", "")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(wd+"\\newDatasetConfig.json", file, 0644)
	if err != nil {
		panic(err)
	}

}

package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

var infoLogger *log.Logger

const (
	host     = "cdm-metadata-test-sg-c.postgres.database.azure.com"
	port     = 5432
	user     = "citus"
	password = "Quickl0ad"
	dbname   = "citus"
)

//TempChar table represents table structure
type TempChar struct {
	versionDatasetID         int64
	dimensionCode            string
	charCode                 string
	charValueCode            string
	charValueDisplayName     string
	charValueDisplaySequence int64
	insertTime               *time.Time
	updateTime               *time.Time
}

func init() {
	logFileName := "C:\\citus\\log\\log.txt"
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	infoLogger = log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.LUTC|log.Lmicroseconds|log.Lshortfile)
}

func citusTest(run int, csvfile string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=verify-full",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	file, err := os.Open(csvfile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	/******************************
	rows, err := db.Query(`select * from metadata.charvalues_temp where version_dataset_id =$1`, 9876)
	if err != nil {
		panic(err)
	}
	tempChars := make([]TempChar, 0)

	for rows.Next() {
		var tempChar TempChar
		if err := rows.Scan(&tempChar.versionDatasetID, &tempChar.dimensionCode, &tempChar.charCode, &tempChar.charValueCode, &tempChar.charValueDisplayName,
			&tempChar.charValueDisplaySequence, &tempChar.insertTime, &tempChar.updateTime); err != nil {
			log.Fatal(err)
		}
		tempChars = append(tempChars, tempChar)
	}
	fmt.Printf("output: %v\n", tempChars)
	*******************************/

	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	stmt, err := tx.Prepare(pq.CopyInSchema("metadata", "charvalues_temp", "version_dataset_id", "dimension_code", "char_code", "char_value_code", "char_value_display_name",
		"char_value_display_sequence", "insert_time", "update_time"))

	if err != nil {
		fmt.Println("in Prepare")
		panic(err)
	}

	startTime := time.Now()
	infoLogger.Println("Begining File Load. Run", run)
	for {
		csvRecord, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if csvRecord[0] == "version_dataset_id" {
			continue
		}

		datasetID, err := strconv.ParseInt(csvRecord[0], 10, 32)
		if err != nil {
			panic(err)
		}
		insertTime := time.Now()
		updateTime := time.Now()

		_, err = stmt.Exec(datasetID, csvRecord[1], csvRecord[2], csvRecord[3], csvRecord[4], csvRecord[5], insertTime, updateTime)
		if err != nil {
			panic(err)
		}

	}

	/*************************************************
	for _, char := range tempChars {
		_, err := stmt.Exec(char.versionDatasetID, char.dimensionCode, char.charCode, char.charValueCode, char.charValueDisplayName, char.charValueDisplaySequence, char.insertTime, char.updateTime)
		if err != nil {
			fmt.Println("in Exec")
			panic(err)
		}
	}
	*************************************************/
	execStartTime := time.Now()

	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}

	endTime := time.Now()
	infoLogger.Println("Completed File Load. Run", run)
	duration := endTime.Sub(startTime)
	execDuration := endTime.Sub(execStartTime)
	infoLogger.Println("Duration for run", run, "is:", duration, "CopyIn duration is:", execDuration)

	err = stmt.Close()
	if err != nil {
		panic(err)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	wg.Done()
}

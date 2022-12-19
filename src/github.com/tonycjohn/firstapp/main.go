package main

import (
	"flag"
	"fmt"
)

//pullrequest test
// CitusConfig represents config
type CitusConfig struct {
	Concurrency int      `json:"concurrency"`
	InputCsvs   []string `json:"inputCsvs"`
	NewVal      string   `json:"newval"`
}

//command line flags
var (
	projectID      = flag.String("project_id", "", "project id")
	snapshotMillis = flag.Int64("snapshot_millis", 0, "snapshot millis")
)

func main() {
	//flag.Parse()
	//fmt.Println("projectID", *projectID)
	//fmt.Println("snapshotMillis", *snapshotMillis)

	fourSum([]int{1, 2, 5, 1, 0, 2}, 9)
	fmt.Println(longestWord("go Tony go"))
	fmt.Println(sortyByParity([]int{1, 3, 4, 9, 7, 6, 12}))

	//publishtoGcp()
	//sendMessageToQue()
	//readFromQue()
	//sendMessageToTopic()
	//readFromTopic()

	//stringReplace(true, false)

	/*startTime := time.Now()
	writeCsv(130000, "12347", "C:\\citus\\csv\\dataset_5\\newbigger130k.csv")
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	log.Println("File Created in ", duration)*/

	/*var config CitusConfig

	configFile, err := os.Open("C:\\goprojects\\src\\github.com\\tonycjohn\\firstapp\\citusConfig.json")
	if err != nil {
		panic(err)
	}

	defer configFile.Close()
	//byteValue, _ := ioutil.ReadAll(configFile)
	//json.Unmarshal(byteValue, &config)

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}

	//fmt.Println("Concurrency-", config.Concurrency, " Length of NewVal-", len(config.NewVal))

	for i, inputCsv := range config.InputCsvs {
		wg.Add(1)
		go citusTest(i, inputCsv)
		fmt.Println(inputCsv)
	}
	wg.Wait()*/

	//playground()
	//webApp()
	//copyFile("C:\\Users\\johnto01\\Downloads\\goTony.docx", "C:\\Users\\johnto01\\Downloads\\tonycjohn.docx")
	//multiThread()
	//readCsv("C:\\Users\\johnto01\\Downloads\\test.csv")

	//numberList() //calling a function

	//httpCall() //calling http request

}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/xuri/excelize/v2"
)

type DataExchange struct {
	Name string `json:"name"`
}

var (
	NameList []string
)

func main() {
	timeStart := time.Now()
	f := excelize.NewFile()
	var data []DataExchange
	for i := 1; i < 5+1; {
		url := fmt.Sprintf("https://api.coingecko.com/api/v3/exchanges?per_page=100&page=%d", i)
		resp, err := http.Get(url)
		if err != nil {
			log.Println(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}

		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Println("The maximum number of requests per minute was exceeded. Performing blocking bypass...")
			time.Sleep(70 * time.Second)
		}
		fmt.Printf("Visited page - %d of %d\n", i, 5)

		for _, val := range data {
			NameList = append(NameList, val.Name)
		}

		for idx, val := range NameList {
			f.SetCellValue("Sheet1", fmt.Sprintf("A%d", idx+1), val)
		}
		i++
	}
	if err := f.SaveAs("ExchangeList.xlsx"); err != nil {
		log.Fatalf("Excel file creation error - %s", err)
	}
	fmt.Println("Total execution time - ", time.Since(timeStart))
}

package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"text/template"
)

type stockMarket struct {
	Date, Open, High, Low, Close, Volume, AdjClose string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	var market []stockMarket
	in := "table.csv"
	csvFile, err := os.Open(in)
	if err != nil {
		log.Fatalln(err)
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		s := stockMarket{
			Date:     record[0],
			Open:     record[1],
			High:     record[2],
			Low:      record[3],
			Close:    record[4],
			Volume:   record[5],
			AdjClose: record[6],
		}

		market = append(market, s)
	}

	err = tpl.Execute(os.Stdout, market)
	if err != nil {
		log.Fatalln(err)
	}
}

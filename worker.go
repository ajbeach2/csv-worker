package csvworker

import (
	"encoding/csv"
	"log"
	"os"
)

type CsvWorker struct {
	File    *os.File
	Writer  *csv.Writer
	In      chan [][]string
	closing chan chan error
}

func (c *CsvWorker) run() {
	for {
		select {
		case errc := <-c.closing:
			errc <- c.File.Close()
			return
		case rows := <-c.In:
			c.Writer.WriteAll(rows)
		}
	}
}

func (c *CsvWorker) Close() error {
	errc := make(chan error)
	c.closing <- errc
	return <-errc

}

func (c *CsvWorker) Recieve(str [][]string) {
	c.In <- str
}

func NewCsvWorker(fileName string) *CsvWorker {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	w := csv.NewWriter(f)

	csvWorker := &CsvWorker{
		File:    f,
		In:      make(chan [][]string),
		Writer:  w,
		closing: make(chan chan error),
	}

	go csvWorker.run()
	return csvWorker
}

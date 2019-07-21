package main

import (
	"csv-import/DataImporter"
	"fmt"
	"log"
	"runtime"
)

func main() {
	PrintMemUsage()

	db, err := DataImporter.NewDb()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", db.Connection.Stats())

	//csvFileReader := filereader.NewCsvFileReader("test/test.csv")
	//
	//count, err := csvFileReader.Count()
	//if err != nil {
	//	fmt.Print(err)
	//}
	//fmt.Printf("\nrow count %v\n", count)
	//
	//var records [][]string
	//for {
	//	record, err := csvFileReader.GetNextDataSet()
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	records = append(records, record)
	//}
	//if err != nil {
	//	fmt.Print(err)
	//}
	//
	//fmt.Printf("\nrecords length %v\n", len(records))

	PrintMemUsage()
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	astisub "github.com/asticode/go-astisub"
)

func main() {
	inputFilePath := flag.String("i", "", "input file (e.g. .vtt)")
	outputFilePath := flag.String("o", *inputFilePath+".txt", "output file (e.g. res.txt)")

	flag.Parse()

	if *inputFilePath == "" {
		log.Fatal("missing required argument: inputFilePath")
	}

	doProcessing(*inputFilePath, *outputFilePath)
}

func doProcessing(inputFilePath, outputFilePath string) {
	start := time.Now()

	fmt.Println("input file:", inputFilePath)
	fmt.Println("output file:", outputFilePath)

	data, err := os.Open(inputFilePath)
	check(err)
	defer data.Close()

	subs, err := astisub.ReadFromWebVTT(data) // FIXME: Check actual file format
	check(err)

	outputFile, err := os.Create(outputFilePath)
	check(err)
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	count := 0

	for _, item := range subs.Items {
		line := formatTimestamp(item.StartAt) + " " + item.String() + "\n"
		_, err := writer.WriteString(line)
		check(err)
		writer.Flush()
		count++
	}

	elapsed := time.Since(start)

	fmt.Printf("finished processing %d items in %f seconds\n", count, elapsed.Seconds())
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func formatTimestamp(duration time.Duration) string {
	hours := int64(duration.Hours())
	duration -= time.Duration(hours) * time.Hour
	minutes := int64(duration.Minutes())
	duration -= time.Duration(minutes) * time.Minute
	seconds := int64(duration.Seconds())
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	astisub "github.com/asticode/go-astisub"
)

const defaultMinTokenCount = 3

func main() {
	inputFilePath := flag.String("i", "", "input file (e.g. .vtt)")
	outputFilePath := flag.String("o", *inputFilePath+".txt", "output file (e.g. res.txt)")

	minTokenCount := flag.Int("minTokenCount", defaultMinTokenCount, "minimum number of tokens (words) for a timestamp line to be registered individually")

	flag.Parse()

	if *inputFilePath == "" {
		log.Fatal("missing required argument: inputFilePath")
	}

	doProcessing(inputFilePath, outputFilePath, minTokenCount)
}

func doProcessing(inputFilePath, outputFilePath *string, minTokenCount *int) {
	start := time.Now()

	fmt.Println("input file:", *inputFilePath)
	fmt.Println("output file:", *outputFilePath)

	data, err := os.Open(*inputFilePath)
	check(err)
	defer data.Close()

	subs, err := astisub.ReadFromWebVTT(data) // FIXME: Check actual file format
	check(err)

	outputFile, err := os.Create(*outputFilePath)
	check(err)
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	var initialCount, finalCount int

	var subText, timestamp string

	// Iterate over subs and write to output file
	for _, item := range subs.Items {
		if timestamp == "" {
			timestamp = formatTimestamp(item.StartAt)
		}

		subText += " " + item.String()
		tokenCount := len(strings.Fields(subText))

		if tokenCount >= *minTokenCount {
			// Dump it
			line := timestamp + " " + subText + "\n"

			_, err := writer.WriteString(line)
			check(err)
			writer.Flush()

			timestamp = ""
			subText = ""
			finalCount++
		}

		initialCount++
	}

	elapsed := time.Since(start)

	fmt.Printf("finished processing in %f seconds\nshrank entry count from %d --> %d\n",
		elapsed.Seconds(), initialCount, finalCount)
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

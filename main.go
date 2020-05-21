package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/codahale/hdrhistogram"
)

// TODO: support reading from file

func main() {
	m := make(map[string]*hdrhistogram.Histogram)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			// TODO: support any whitespace here
			// TODO: warn on bad lines instead of erroring
			// TODO: skip empty lines
			// TODO: custom separator to support keys containing spaces (or just
			//       use the last space of the line to split)
			log.Fatalf("bad log line, expected k <space> value, got %v", parts)
		}

		key := parts[0]
		rawval := parts[1]
		val, err := strconv.ParseInt(rawval, 10, 64)
		if err != nil {
			log.Fatalf("bad value, expected int, got %v", rawval)
		}

		hdr, ok := m[key]
		if !ok {
			// TODO: make all of these parameters configurable.
			//       alternatively we can use linear and exponential
			//       buckets that can grow dynamically.
			// minValue, maxValue int64, sigfigs int
			hdr = hdrhistogram.New(10, 1000, 2)
			m[key] = hdr
		}

		hdr.RecordValue(val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// TODO: sort keys (by total?)
	// TODO: top-k mode for high-cardinality keys

	for k, hdr := range m {
		largestBucket := int64(0)
		for _, bar := range hdr.Distribution() {
			if bar.Count > largestBucket {
				largestBucket = bar.Count
			}
		}

		numDigits := len(strconv.Itoa(int(largestBucket)))
		numFormat := fmt.Sprintf("%%%dd", numDigits)

		fmt.Printf("%s (%d)\n", k, hdr.TotalCount())
		for _, bar := range hdr.Distribution() {
			// TODO: collapse buckets (?)
			// if bar.Count == 0 {
			// 	continue
			// }

			// TODO: make scaleFactor configurable
			// TODO: indent everything with spaces (proper tables)
			// TODO: figure out the proper scientific notation for inclusive/exclusive range

			pct := int(bar.Count * 100 / largestBucket)
			scaleFactor := 0.4
			fmt.Printf(
				"[%d,\t%d]\t[%s%s]\t(%s/%d)\n",
				bar.From,
				bar.To,
				strings.Repeat("=", int(math.Floor(float64(pct)*scaleFactor))),
				strings.Repeat(" ", int(math.Ceil(float64(100-pct)*scaleFactor))),
				fmt.Sprintf(numFormat, bar.Count),
				hdr.TotalCount(),
			)
		}
		fmt.Println()
	}
}

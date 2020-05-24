package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/igorwwwwwwwwwwwwwwwwwwww/log2hist/hist"
	"github.com/pkg/profile"
	"github.com/titanous/bconv"
)

// TODO: scaling factor (e.g. bytes => GiB)

var group = flag.Bool("g", false, `if enabled, input is expected in "key val" format; defaults to false`)
var pprof = flag.Bool("pprof", false, `if enabled, cpu profile is taken; defaults to false`)

func main() {
	if *pprof {
		defer profile.Start().Stop()
	}

	var err error

	m := make(map[string]*hist.Histogram)

	flag.Parse()
	args := flag.Args()

	r := os.Stdin
	if len(args) > 0 {
		file := args[0]
		r, err = os.Open(file)
		if err != nil {
			log.Fatalf("could not open file %v: %v", file, err)
		}
	}

	scanner := bufio.NewScanner(bufio.NewReader(r))
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		key := ""
		rawval := line

		if *group {
			fields := strings.Fields(line)
			if len(fields) != 2 {
				log.Printf("warning: ignoring bad value, expected k <space> value, got %v", fields)
				continue
			}

			key = fields[0]
			rawval = fields[1]
		}

		val, err := strconv.ParseUint(rawval, 10, 64)
		if err != nil {
			log.Printf("warning: ignoring bad value, expected int, got %v", rawval)
			continue
		}

		h, ok := m[key]
		if !ok {
			h = hist.New()
			m[key] = h
		}

		err = h.Record(val)
		if err != nil {
			log.Printf("warning: ignoring bad value %v, got error: %v", val, err)
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if !*group {
		// in non group-by mode, the empty key is used to store the global histogram
		fmt.Println(m[""])
		return
	}

	// TODO: top-k mode for high-cardinality keys

	// sort keys by total count
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		ki := keys[i]
		kj := keys[j]
		mi := m[ki]
		mj := m[kj]
		if mi.Count == mj.Count {
			return ki > kj
		}
		return mi.Count > mj.Count
	})

	for _, k := range keys {
		fmt.Println(k)
		fmt.Println(m[k])
	}
}

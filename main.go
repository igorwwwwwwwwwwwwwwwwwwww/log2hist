package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
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

func run(r io.Reader, w io.Writer) error {
	h := hist.New()

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Bytes()

		if len(line) == 0 {
			continue
		}

		rawval := line

		val, err := bconv.ParseUint(rawval, 10, 64)
		if err != nil {
			log.Printf("warning: ignoring bad value, expected int, got %s", rawval)
			continue
		}

		err = h.Record(val)
		if err != nil {
			log.Printf("warning: ignoring bad value %v, got error: %v", val, err)
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Fprintln(w, h)

	return nil
}

func runWithGroup(r io.Reader, w io.Writer) error {
	m := make(map[string]*hist.Histogram)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Bytes()

		if len(line) == 0 {
			continue
		}

		fields := bytes.Fields(line)
		if len(fields) != 2 {
			log.Printf("warning: ignoring bad value, expected k <space> value, got %s", fields)
			continue
		}

		key := fields[0]
		rawval := fields[1]

		val, err := bconv.ParseUint(rawval, 10, 64)
		if err != nil {
			log.Printf("warning: ignoring bad value, expected int, got %s", rawval)
			continue
		}

		strkey := string(key)
		h, ok := m[strkey]
		if !ok {
			h = hist.New()
			m[strkey] = h
		}

		err = h.Record(val)
		if err != nil {
			log.Printf("warning: ignoring bad value %v, got error: %v", val, err)
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		return err
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
		fmt.Fprintln(w, k)
		fmt.Fprintln(w, m[k])
	}

	return nil
}

func main() {
	if *pprof {
		defer profile.Start().Stop()
	}

	var err error

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

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	if !*group {
		err = run(r, w)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err = runWithGroup(r, w)
	if err != nil {
		log.Fatal(err)
	}
}

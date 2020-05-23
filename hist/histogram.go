package hist

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var RecordError = errors.New("value must be within range 0..2^63")

type Histogram struct {
	Bins  []uint64
	Count uint64
}

func New() *Histogram {
	return &Histogram{
		Bins:  make([]uint64, 64),
		Count: 0,
	}
}

func indexLabel(power int) string {
	return strconv.Itoa(1 << power)
}

// https://github.com/iovisor/bpftrace/blob/1ece0d0b1441aa70d4a6b324fb852954a5989eab/src/output.cpp#L166
func (h Histogram) String() string {
	var minIdx, maxIdx int
	var maxVal uint64

	for i, val := range h.Bins {
		if val == 0 {
			continue
		}
		if i < minIdx || (minIdx == 0 && h.Bins[0] == 0) {
			minIdx = i
		}
		if i > maxIdx {
			maxIdx = i
		}
		if val > uint64(maxVal) {
			maxVal = val
		}
	}

	var out string

	for i := minIdx; i <= maxIdx; i++ {
		var header string
		if i == 0 {
			header = "[0]"
		} else if i == 1 {
			header = "[1]"
		} else {
			header = "[" + indexLabel(i-1) + ", " + indexLabel(i) + ")"
		}

		maxWidth := 52
		barWidth := uint64((float64(h.Bins[i]) / float64(maxVal)) * float64(maxWidth))
		bar := strings.Repeat("@", int(barWidth))

		out = out + fmt.Sprintf("%16s %8d |%-52s|\n", header, h.Bins[i], bar)
	}

	return out
}

// TODO: fix 1 being recorded as bucket 0

func (h *Histogram) Record(val uint64) error {
	i := uint64(math.Ceil(1 + math.Log2(float64(val))))
	if i >= 64 || i < 0 {
		return RecordError
	}
	h.Bins[i]++
	h.Count++
	return nil
}

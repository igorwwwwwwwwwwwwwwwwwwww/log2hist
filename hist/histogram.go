package hist

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var ErrRecord = errors.New("value must be within range 0..2^63")

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
	var minIdx = -1
	var maxIdx = 0
	var maxVal uint64

	for i, val := range h.Bins {
		if val > 0 {
			if minIdx == -1 {
				minIdx = i
			}
			maxIdx = i
		}
		if val > maxVal {
			maxVal = val
		}
	}

	if minIdx == -1 {
		return ""
	}

	var out string

	for i := minIdx; i <= maxIdx; i++ {
		var header string
		if i == 0 {
			header = "(..., 0)"
		} else if i == 1 {
			header = "[0]"
		} else if i == 2 {
			header = "[1]"
		} else {
			header = "[" + indexLabel(i-2) + ", " + indexLabel(i-1) + ")"
		}

		maxWidth := 52
		barWidth := uint64((float64(h.Bins[i]) / float64(maxVal)) * float64(maxWidth))
		bar := strings.Repeat("∎", int(barWidth))

		out = out + fmt.Sprintf("%16s %8d  %-52s\n", header, h.Bins[i], bar)
	}

	return out
}

func (h *Histogram) Record(val uint64) error {
	i := uint64(math.Floor(2 + math.Log2(float64(val))))
	if val == 0 {
		i = 1
	}
	if i >= 64 {
		return ErrRecord
	}
	h.Bins[i]++
	h.Count++
	return nil
}

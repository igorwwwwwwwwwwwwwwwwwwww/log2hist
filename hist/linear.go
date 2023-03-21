package hist

import (
	"fmt"
	"strconv"
	"strings"
)

type LinearHistogram struct {
	Min, Max, Step int

	Bins  []uint64
	Count uint64
}

func NewLinear(min, max, step int) Histogram {
	buckets := (max - min) / step

	return &LinearHistogram{
		Min:   min,
		Max:   max,
		Step:  step,
		Bins:  make([]uint64, buckets+2),
		Count: 0,
	}
}

func (h LinearHistogram) indexLabel(number int) string {
	return strconv.Itoa(number)
}

// https://github.com/iovisor/bpftrace/blob/1ece0d0b1441aa70d4a6b324fb852954a5989eab/src/output.cpp#L205
func (h LinearHistogram) String() string {
	var maxIdx = -1
	var maxVal uint64

	for i, val := range h.Bins {
		if val != 0 {
			maxIdx = i
		}
		if val > maxVal {
			maxVal = val
		}
	}

	if maxIdx == -1 {
		return ""
	}

	var startVal, endVal int

	buckets := (h.Max - h.Min) / h.Step
	for i := 0; i <= buckets+1; i++ {
		val := h.Bins[i]
		if val > 0 {
			if startVal == -1 {
				startVal = i
			}
			endVal = i
		}
	}

	if startVal == -1 {
		startVal = 0
	}

	var out string

	for i := startVal; i <= endVal; i++ {
		var header string
		if i == 0 {
			header = "(..., " + h.indexLabel(h.Min) + ")"
		} else if i == int(buckets+1) {
			header = "[" + h.indexLabel(h.Max) + ", ...)"
		} else {
			header = "[" + h.indexLabel((i-1)*h.Step+h.Min) + ", " + h.indexLabel(i*h.Step+h.Min) + ")"
		}

		maxWidth := 52
		barWidth := uint64((float64(h.Bins[i]) / float64(maxVal)) * float64(maxWidth))
		bar := strings.Repeat("∎", int(barWidth))

		out = out + fmt.Sprintf("%16s %8d  %-52s\n", header, h.Bins[i], bar)
	}

	return out
}

// https://github.com/iovisor/bpftrace/blob/70ee22cb14e2eedc5df17e53965824d7381f8e6f/src/ast/passes/codegen_llvm.cpp#L2980-L2991
func (h *LinearHistogram) Record(val uint64) error {
	i := uint64(h.index(val))
	h.Bins[i]++
	h.Count++
	return nil
}

func (h LinearHistogram) index(val uint64) int {
	if int(val) < h.Min {
		return 0
	}
	if int(val) > h.Max {
		return 1 + (h.Max-h.Min)/h.Step
	}
	return 1 + (int(val)-h.Min)/h.Step
}

func (h LinearHistogram) GetCount() uint64 {
	return h.Count
}

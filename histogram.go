package histogram

import (
	"sort"
	"math"
)

type Bins []*Bin

type Bin struct {
	Point float64
	Count float64
}

type Histogram struct {
	Size int
	Bins Bins
}

func New(n int) *Histogram {
	return &Histogram{ Size: n }
}

func (b Bins) Len() int {
	return len(b)
}

func (b Bins) Less(i, j int) bool {
	return b[i].Point < b[j].Point
}

func (b Bins) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (h *Histogram) Insert(n float64) {
	h.update(n)
}

func (h *Histogram) Merge(b Bins) {
	h.Bins = append(h.Bins, b...)
	sort.Sort(h.Bins)
	h.compress()
}

// private

func (h *Histogram) update(n float64) {
	for i := 0; i < len(h.Bins); i++ {
		if h.Bins[i].Point == n {
			h.Bins[i].Count++
			return
		}
	}

	h.Bins = append(h.Bins, &Bin{n, 1})
	sort.Sort(h.Bins)

	if len(h.Bins) > h.Size {
		h.compress()
	}
}

func (h *Histogram) compress() {
	for len(h.Bins) > h.Size {
		idx := -1
		val := math.MaxFloat64

		for i := 0; i < len(h.Bins)-1; i++ {
			cur := h.Bins[i]
			nxt := h.Bins[i+1]

			if gap := nxt.Point - cur.Point; gap < val {
				idx = i
				val = gap
			}
		}

		cur := h.Bins[idx]
		nxt := h.Bins[idx+1]

		m := cur.Count + nxt.Count
		p := (cur.Point + nxt.Point) / m


		h.Bins = append(h.Bins[:idx], h.Bins[idx+2:]...)
		h.Bins = append(h.Bins, &Bin{p,m})
		sort.Sort(h.Bins)
	}
}

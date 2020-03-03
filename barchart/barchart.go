package barchart

import (
	"errors"
	"math"
)

func PlotBarChart(data []float64, colors []uint32, backgroundColor uint32, h, w int) ([]uint32, error) {
	screen := make([]uint32, (h * w))

	// Make sure that data is well-formed
	if len(data) != len(colors) {
		return nil, errors.New("Size mismatch between data and color arrays.")
	}

	if len(data) == 0 {
		for i := 0; i < len(screen); i++ {
			screen[i] = backgroundColor
		}
		return screen, nil
	}

	max := -(math.MaxFloat64)
	for i := range data {
		if data[i] > max {
			max = data[i]
		}
	}

	for i := range data {
		data[i] = (data[i] / max) * 0.875
	}

	barWidth := w / len(data)
	index := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			barIx := j / barWidth
			if (barIx < len(data)) && (i < int(data[barIx]*float64(h))) {
				screen[index] = colors[barIx]
			} else {
				screen[index] = backgroundColor
			}
			index++
		}
	}

	return screen, nil
}

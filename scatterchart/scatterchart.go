package scatterchart


import(
  "errors"
  "math"
)





func PlotScatterChart(xs, ys []float64, colors []uint32, backgroundColor uint32, h, w, radius int) ([]uint32, error) {
  screen := make([]uint32, (h * w), backgroundColor);

  // Make sure that data is well-formed
  if (len(xs) != len(ys)){
    return nil, errors.New("Size mismatch between x and y coordinate arrays.")
  }
  if (len(xs) != len(colors)){
    return nil, errors.New("Size mismatch between data and color arrays.")
  }

  if (len(xs) == 0){
    for i := 0; i < len(screen); i++ {
      screen[i] = backgroundColor
    }
    return screen, nil
  }

  maxX := -(math.MaxFloat64)
  minX :=  (math.MaxFloat64)
  maxY := -(math.MaxFloat64)
  minY :=  (math.MaxFloat64)
  for i := range(xs) {
    if xs[i] > maxX {
      maxX = xs[i]
    }
    if xs[i] < minX {
      minX = xs[i]
    }
    if ys[i] > maxY {
      maxY = ys[i]
    }
    if ys[i] < minY {
      minY = ys[i]
    }
  }

  dx := maxX - minX
  dy := maxY - minY

  for i := range(xs) {
    xs[i] = 0.0625 + (((xs[i] - minX) / dx) * 0.875)
    ys[i] = 0.0625 + (((ys[i] - minY) / dy) * 0.875)
  }


  for i := range(xs) {
    x := int(xs[i] * float64(w))
    y := int(ys[i] * float64(h))

    for j := 0; j < radius; j++ {
      for k := 0; k < radius; k++ {
        px := x + j - (radius / 2)
        py := y + k - (radius / 2)

        screen[(py * w) + px] = colors[i]
      }
    }
  }



  return screen, nil
}

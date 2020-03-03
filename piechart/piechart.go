package piechart






func PlotPieChart(data []float64, colors []uint32, backgroundColor uint32, h, w int) []uint32 {
  screen := make([]uint32, (h * w));

  index := 0
  for i := 0; i < h; i++{
    x := (2.0 * (float64(i) / float64(h))) - 1.0
    for j := 0; j < w; j++{
      y := (2.0 * (float64(j) / float64(w))) - 1.0
      if (x*x) + (y*y) <= 0.8 {
        screen[index] = colors[0]
      }else{
        screen[index] = backgroundColor
      }
      index++
    }
  }

  return screen
}

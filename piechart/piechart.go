package piechart


import(
  "errors"
  "sort"
  "math"
  "fmt"
)





func PlotPieChart(data []float64, colors []uint32, backgroundColor uint32, h, w int) ([]uint32, error) {
  screen := make([]uint32, (h * w));

  // Make sure that data is well-formed
  if (len(data) != len(colors)){
    return nil, errors.New("Size mismatch between data and color arrays.")
  }

  if (len(data) == 0){
    for i := 0; i < len(screen); i++ {
      screen[i] = backgroundColor
    }
    return screen, nil
  }

  // Make a table so that things are searchable
  sizes := make([]float64, len(data))
  for i := range(data) {
    if i != 0 {
      sizes[i] = data[i] + sizes[i-1]
    }else{
      sizes[i] = data[i]
    }
  }
  scale := sizes[len(sizes)-1]
  for i := range(data){
    sizes[i] /= scale
  }

  fmt.Println(sizes)

  index := 0
  min   :=  10.0
  max   := -10.0
  for i := 0; i < h; i++{
    x := (2.0 * (float64(i) / float64(h))) - 1.0
    for j := 0; j < w; j++{
      y := (2.0 * (float64(j) / float64(w))) - 1.0
      if (x*x) + (y*y) <= 0.875 {
        // Figure out which color to draw here
        angle := 0.5 + (math.Atan2(x, y) / (2 * math.Pi))
        //fmt.Println(angle)
        if angle > max {
          max = angle
        }
        if angle < min {
          min = angle
        }

        colorIx := sort.SearchFloat64s(sizes, angle)

        if(colorIx >= 0) && (colorIx < len(colors)){
          screen[index] = colors[colorIx]
        }

      }else{
        screen[index] = backgroundColor
      }
      index++
    }
  }

  fmt.Println(min, max)

  return screen, nil
}

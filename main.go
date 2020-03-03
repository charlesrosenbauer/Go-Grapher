package main

import(
  "github.com/charlesrosenbauer/Go-Grapher/piechart"
  "github.com/veandco/go-sdl2/sdl"
)




func main(){

  // Set up the screen for SDL
  var window *sdl.Window
  var screen *sdl.Surface
  if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil{
    panic(err)
  }
  if w, err := sdl.CreateWindow("Go-grapher", 0, 0, 512, 512, 0); err != nil{
    panic(err)
  }else{
    window = w
  }
  if s, err := window.GetSurface(); err != nil {
    panic(err)
  }else{
    screen = s
  }

  chart := piechart.PlotPieChart(nil, []uint32{0xff,0xff00}, 0x0f0f0f, 256, 256)

  pixels := screen.Pixels()

  chartIndex := 0
  for i := 0; i < 256; i++ {
    for j := 0; j < 256; j++ {
      pixelIndex := (i * 4) + (j * 512 * 4)
      pixels[pixelIndex  ] = byte((chart[chartIndex] >>  0) & 255)
      pixels[pixelIndex+1] = byte((chart[chartIndex] >>  8) & 255)
      pixels[pixelIndex+2] = byte((chart[chartIndex] >> 16) & 255)
      pixels[pixelIndex+3] = byte((chart[chartIndex] >> 24) & 255)
      chartIndex++
    }
  }

  window.UpdateSurface()

  sdl.Delay(1000)


}

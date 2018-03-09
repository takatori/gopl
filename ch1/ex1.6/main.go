package main

import (
	"flag"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
	color.Black,
	color.White,
	color.RGBA{0x00, 0x80, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
	color.RGBA{0xff, 0x00, 0x00, 0xff},
}

const (
	blackIndex = 0 // パレットの次の色
	whiteIndex = 1 // パレットの最初の色
	greenIndex = 2
	blueIndex  = 3
	redIndex   = 4
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	i := flag.Int("c", 1, "color index")
	flag.Parse()
	lissajous(os.Stdout, uint8(*i))
}

func lissajous(out io.Writer, colorIndex uint8) {
	const (
		cycles  = 5     // 発振器xが完了する周回の回数
		res     = 0.001 // 回転の分解能
		size    = 100   //画像キャンパスは[-size..+size]の範囲を扱う
		nframes = 64    // アニメーションフレーム数
		delay   = 8     //10ms単位でのフレーム間遅延
	)
	freq := rand.Float64() * 3.0 // 発振器yの相対周波数
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //位相差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

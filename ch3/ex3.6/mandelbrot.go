// mandelbrotはマンデルブロフラクタルのPNG画像を生成します
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 2048, 2048
	//width, height          = 1024, 1024
)

func main() {

	img := image.NewRGBA(image.Rect(0, 0, width/2, height/2))

	for py := 0; py < height-2; py += 2 {

		y1 := float64(py)/height*(ymax-ymin) + ymin
		y2 := float64(py+1)/height*(ymax-ymin) + ymin

		for px := 0; px < width-2; px += 2 {

			x1 := float64(px)/width*(xmax-xmin) + xmin
			x2 := float64(px+1)/width*(xmax-xmin) + xmin

			z1 := complex(x1, y1)
			z2 := complex(x2, y1)
			z3 := complex(x1, y2)
			z4 := complex(x2, y2)
			// 画像の点(px, py)は複素数値zを表している

			img.Set(px/2, py/2, mandelbrot((z1+z2+z3+z4)/4))
		}
	}

	png.Encode(os.Stdout, img) // 注意: エラーを無視
}

func mandelbrot(z complex128) color.Color {

	const contrast = 15
	const iterations = 200 / contrast

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

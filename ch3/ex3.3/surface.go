package main

import (
	"fmt"
	"image/color"
	"math"
)

const (
	width, height = 600, 320            // キャンバスの大きさ(画素数)
	cells         = 100                 // 格子のます目の数
	xyrange       = 30.0                // 軸の範囲(-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // x単位及びy単位あたりの画素数
	zscale        = height * 0.4        // z単位あたりの画素数
	angle         = math.Pi / 6         // x, y軸の角度(=30度)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30度), cos(30度)
var zmax, zmin = 0.0, 100.0                         // 高さの最大・最小値
var p = make([][]float64, cells*cells)              // 変換後のポリゴンの座標

func main() {

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)
			z := (az + bz + cz + dz) / 4
			if z > zmax {
				zmax = z
			}
			if z < zmin {
				zmin = z
			}
			p[i*cells+j] = []float64{ax, ay, bx, by, cx, cy, dx, dy, z}
		}
	}
    
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' style='stroke: grey; fill: white; stroke-width: 0.7' width='%d' height='%d'>", width, height)
	for i := 0; i < cells*cells; i++ {       
		r := uint8((p[i][8] - zmin)/ (zmax - zmin) * 255)
		b := 255 - r
		c := color.RGBA{r, 0, b, 0}
		fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' style='fill:%v' />\n",
			p[i][0], p[i][1], p[i][2], p[i][3], p[i][4], p[i][5], p[i][6], p[i][7], rgbToString(c))
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	// ます目(i, j)のかどの点(x, y)を見つける。
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// 面の高さzを計算する。
	z := f(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // (0, 0)からの距離
	return math.Sin(r) / r
}

func rgbToString(c color.RGBA) string {
	return fmt.Sprintf("rgb(%d,%d,%d)", c.R, c.G, c.B)
}

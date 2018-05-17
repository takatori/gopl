package surface

import (
	"fmt"
	"math"
)

const (
	cells = 100         // 格子のます目の数
	angle = math.Pi / 6 // x, y軸の角度(=30度)
)

var width, height = 600, 320             // キャンバスの大きさ(画素数)
var xyrange = 30.0                       // 軸の範囲(-xyrange..+xyrange)
var xyscale = float64(width/2) / xyrange // x単位及びy単位あたりの画素数
var zscale = float64(height) * 0.4       // z単位あたりの画素数

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30度), cos(30度)

func initCanvas(h int, w int) {
	width = w
	height = h
	xyscale = float64(width/2) / xyrange
	zscale = float64(height) * 0.4
}

// SVG is
func SVG(h int, w int, c string) string {

    fmt.Printf("%d, %d, %s\n", h, w, c)

    initCanvas(h, w)

	s := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' style='stroke: grey; fill: %s; stroke-width: 0.7' width='%d' height='%d'>", c, width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			s += fmt.Sprintf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	s += fmt.Sprintln("</svg>")
	return s
}

func corner(i, j int) (float64, float64) {

	// ます目(i, j)のかどの点(x, y)を見つける。
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 面の高さzを計算する。
	z := f(x, y)

	// (x, y, z)を2-D SVGキャンバス(sx,sy)へ等角的に投影
	sx := float64(width/2) + (x-y)*cos30*xyscale
	sy := float64(height/2) + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // (0, 0)からの距離
	return math.Sin(r) / r
}

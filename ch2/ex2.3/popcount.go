package popcount

// pc[i]はiのポピュレーションカウント
var pc [256]byte

// プログラムが開始した時点で宣言されてした順序で自動的に実行される
// 呼び出すことも参照することもできない
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount はXのポピュレーションカウントを返す
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCount はXのポピュレーションカウントを返す
func PopCountLoop(x uint64) int {

	var num byte
	var i uint64
	for i = 0; i < 8; i++ {
		num += pc[byte(x>>(i*8))]
	}
	return int(num)
}

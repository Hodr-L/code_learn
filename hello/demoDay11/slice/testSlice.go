package slice

func TestSlice(n int) []uint64 {
	/*
		因为要返回斐波那契额数列的切片，
		所以返回值为uint64,
		uint64是存放最大值的数据类型
	*/
	a := make([]uint64, n)
	a[0] = 1
	a[1] = 1
	for i := 2; i < n; i++ {
		//for循环存放斐波那契额数列
		a[i] = a[i-1] + a[i-2]
	}
	return a
}

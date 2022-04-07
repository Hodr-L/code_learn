package sum

func Addupper() func(int) int { //定义Add函数，返回值为匿名函数
	//闭包累加器
	var n int = 10           //初始n值为10，返回n值后n值被返回值替换
	return func(x int) int { //返回匿名函数，匿名函数返回int
		n = n + x
		return n //返回n值
	}
}

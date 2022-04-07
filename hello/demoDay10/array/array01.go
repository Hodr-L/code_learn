package array

import "fmt"

func Array01() {
	var hens [6]float64
	var he float64
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0
	for i := 0; i < len(hens); i++ {
		he += hens[i]
	}
	num := he / float64(len(hens))
	fmt.Printf("%.2f\n", num)
}

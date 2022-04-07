package sum

import "fmt"

func Sum4(i int) { //a int
	for j := 1; j <= i; j++ {
		for n := 1; n <= (2*j)-1; n++ {
			fmt.Print("*")
		}
		fmt.Printf("\n")
	}

}

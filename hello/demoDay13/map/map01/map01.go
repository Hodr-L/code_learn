package map01

import "fmt"

func Map01() {
	//var map02 = map[int]string{
	//	1: "1",
	//	2: "2",
	//}
	//fmt.Println(map02)
	//stu := map[string]map[string]string{
	//	"1": {
	//		"tom": "m",
	//	},
	//	"2": {
	//		"jess": "m",
	//	},
	//	"3": {
	//		"hod": "w",
	//	},
	//	"4": {
	//		"kai": "m",
	//	},
	//	"5": {
	//		"low": "w",
	//	},
	//}
	//val, find := stu["5"]
	//if find {
	//	fmt.Println("找到了：", val)
	//} else {
	//	fmt.Println("没有找到！")
	//}

	sli := make([]map[string]string, 1)
	if sli[0] == nil {
		sli[0] = make(map[string]string, 1)
		sli[0]["names"] = "牛魔王"
		sli[0]["age"] = "500"
	}
	newsli := map[string]string{
		"name": "孙悟空",
		"age":  "800",
	}

	sli = append(sli, newsli)
	fmt.Println(sli)
	for c, v := range sli {
		for k, i := range v {
			fmt.Println(c, k, i)
		}
	}
}

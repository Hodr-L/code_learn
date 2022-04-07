package day19tes01t

import "fmt"

func Big(items ...interface{}) {
	for index, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("第%v个参数是bool类型，值是%v\n", index+1, v)
		case int, int32, int64:
			fmt.Printf("第%v个参数是int类型，值是%v\n", index+1, v)
		case float64, float32:
			fmt.Printf("第%v个参数是float类型，值是%v\n", index+1, v)
		case string:
			fmt.Printf("第%v个参数是string类型，值是%v\n", index+1, v)
		default:
			fmt.Println("不存在该类型！")
		}
	}

}

package chazhao

import "fmt"

func Shunxu(hero string) {
	arr := []string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	for c, i := range arr {
		if hero == i {
			fmt.Println(i)
			break
		} else if c == (len(arr) - 1) {
			fmt.Println("名字不存在！")
		}
	}
}
func Shunxu2(hero string) {
	arr := []string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	index := -1
	for c, v := range arr {
		if hero == v {
			fmt.Println("找到了！", v)
			index = c
			break
		}
	}
	if index == -1 {
		fmt.Println("名字不存在！")
	}
}

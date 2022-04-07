package twoArr

import "fmt"

func TesttwoArr() {
	class := [3][5]float64{}
	total := 0.0
	avg := 0.0
	for i := 0; i < len(class); i++ {
		for j := 0; j < len(class[i]); j++ {
			fmt.Printf("请输入第%v个班级第%v名学生的成绩：", i+1, j+1)
			fmt.Scanln(&class[i][j])
		}
	}
	for j, v := range class {
		and := 0.0
		for _, i := range v {
			and += i
		}
		total += and
		avg = total / float64(len(v)*3)
		fmt.Printf("第%v个班的平均成绩为：%v\n", j+1, and/float64(len(v)))
	}
	fmt.Printf("所有班级的平均成绩为：%v", avg)
}

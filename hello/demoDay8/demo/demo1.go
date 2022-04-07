package demo

import (
	"errors"
	"fmt"
)

func Rune(name string) (err error) {
	//遍历汉字字符串
	//i := []rune(r)
	//for j := 0; j < len(i); j++ {
	//	fmt.Printf("%c\n", i[j])
	//}

	//字符串检验整数 ！并不是字符串类类型转换成整数类型
	//i, err := strconv.Atoi(r)
	//if err != nil {
	//	fmt.Println("错误！！！", err)
	//} else {
	//	fmt.Println(i)
	//}
	//i := string([]byte{r})

	//进制转换
	//i := strconv.FormatInt(r, 16)
	//fmt.Printf("%v", i)

	//i := strings.EqualFold(r, "ABCD")

	//i := strings.Index(r, "abc")

	//i := strings.Replace(r, "a", "b", -1)
	//i := strings.Split(r, ",")
	//fmt.Println(i)

	//defer func() {
	//	err := recover()
	//	if err != nil {
	//		fmt.Println("err", err)
	//		fmt.Println("liuheng@16546.com")
	//	}
	//}()
	//num1 := 10
	//num2 := 0
	//fmt.Println(num1 / num2)

	if name == "config.ini" { //2
		//读取
		return nil //3
	} else { //3
		//返回一个自定义错误
		return errors.New("读取文件错误")
	}

}

func Test02(i string) {
	err := Rune(i)  //1
	if err != nil { //4
		panic(err) //如果读取文件发生错误,就发出错误并终止程序
	}
	fmt.Println("继续执行!")
}

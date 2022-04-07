package file

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func File() {
	//1.带缓冲读取文件
	//file, _ := os.Open("F:/hello.txt")
	////if err != nil {
	////	fmt.Println("打开文件失败！")
	////}
	//defer file.Close() //及时关闭句柄，防止内存泄露
	////创建一个*Reader,带有缓冲
	///*
	//	const (
	//		defaultBufSize = 4096//默认的缓冲大小为4096
	//	)
	//*/
	//reader := bufio.NewReader(file)
	//for {
	//	str, err := reader.ReadString('\n') //读到个换行就结束！
	//	if err == io.EOF {                  //读取到文件末尾
	//		break
	//	}
	//	fmt.Print(str)
	//}
	//fmt.Print("文件读取结束...")
	////if err != nil {
	////	fmt.Println("关闭失败！")
	////}
	//2.一次性读取文件
	file := "f:/hello.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("出错了！")
	} else {
		//fmt.Println(content) //[]byte 因为ioutil.ReadFile函数返回值为byte切片，所以返回byte型数据
		fmt.Println(string(content)) //将[]byte型数据转为string型数据
		//因为我们没有显式的open，所以不需要显式的close
		//因为ioutil.ReadFile函数封装了open和close
	}
}
func File2() {
	filePath := "f:/helloo.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	//os.O_APPEND追加
	//|os.O_CREATE 用于建立文件
	//os.O_TRUNC 清空内容，用于重写
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer file.Close()
	str := "wang\n"
	Write := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		Write.WriteString(str)
	}
	//因为Write是带有缓存的，所以写入内容没有写入磁盘，而是写入缓存中
	//所以需要调用flush方法，将缓存数据写入磁盘中，否则文件中没有数据
	Write.Flush()

	reader, _ := ioutil.ReadFile(filePath)
	fmt.Println(string(reader))

}

func File3() {
	filepath := "F:/hello.txt"
	filepath2 := "F:/helloo.txt"
	data, err := ioutil.ReadFile(filepath2)
	if err != nil {
		fmt.Println("读取错误！")
		return
	} //else {
	//	fmt.Println(string(data))
	//}
	err = ioutil.WriteFile(filepath, data, 0666)
	if err != nil {
		fmt.Println("写入出错！")
	}
}

//常用的判断文件是否存在函数↓
func File4(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		fmt.Println("文件存在！")
		return true, nil
	}
	if os.IsNotExist(err) {
		fmt.Println("文件存在！")
		return false, nil
	}
	return false, nil
}

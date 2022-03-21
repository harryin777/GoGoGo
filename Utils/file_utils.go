package Utils

import (
	"fmt"
	"os"
)

func CreateFile(fileName string) *os.File {
	var f *os.File
	var err error
	//f, err = os.Create(filename)
	//if err != nil {
	//	panic(err)
	//}
	if checkFileIsExist(fileName) { //如果文件存在
		err := os.Remove(fileName)
		if err != nil {
			panic(err)
		}
		fmt.Println("文件存在, 删除")
	} else {
		f, err = os.Create(fileName) //创建文件
		if err != nil {
			panic(err)
		}
		fmt.Println("文件不存在")
	}
	return f
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

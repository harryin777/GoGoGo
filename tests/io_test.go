package tests

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

/*
 	O_RDONLY 以只读文式打开文件。
    O_WRONLY 以只写方式打开文件。
    O_RDWR 以读写方式打开文件
    O_APPEND 以追加方式打开文件，写入的数据将追加到文件尾。
    O_CREATE 当文件不存在时创建文件。
    O_EXCL 与 O_CREATE 一起使用，当文件已经存在时 Open 操作失败。
    O_SYNC 以同步方式打开文件。每次 write 系统调用后都等待实际的物理 I/O 完成后才返回，默认(不使用该标记)是使用缓冲的，也就是说每次的写操作是写到系统内核缓冲区中，等系统缓冲区满后才写到实际存储设备。
    O_TRUNC 如果文件已存在，打开时将会清空文件内容。必须于 O_WRONLY 或 O_RDWR 配合使用。截断文件，需要有写的权限。
*/
func Test_Write(t *testing.T) {
	var filename = "./output1.txt"
	var f *os.File
	var err error
	if checkFileIsExist(filename) { //如果文件存在
		f, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0666) //打开文件
		if err != nil {
			panic(err)
		}
		fmt.Println("文件存在")
	} else {
		f, err = os.Create(filename) //创建文件
		if err != nil {
			panic(err)
		}
		fmt.Println("文件不存在")
	}
	w := bufio.NewWriter(f) //创建新的 Writer 对象
	n4, err := w.WriteString("bufferedn")
	if err != nil {
		panic(err)
	}
	fmt.Printf("写入 %d 个字节n", n4)
	w.Flush()
	f.Close()
}

func TestReader(t *testing.T) {
	buf := bufio.NewReader(os.Stdin)
	for {
		str, err := buf.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Println(str)
	}
}

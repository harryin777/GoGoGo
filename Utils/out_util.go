/**
 * @Author: yky
 * @Date: 2021/5/12 14:56
 * @Version: 1.0
 */
package Utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func ReceiveStr(str string) {
	var strBuffer bytes.Buffer
	_ = json.Indent(&strBuffer, []byte(str), "", "    ")
	fmt.Println("formated: ", strBuffer.String())
}

func ReceiveStruct(t interface{}) {
	b, err := json.Marshal(t)
	if err != nil {
		log.Fatalln(err)
	}

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "\t")

	if err != nil {
		log.Fatalln(err)
	}

	out.WriteTo(os.Stdout)
	fmt.Println()
}

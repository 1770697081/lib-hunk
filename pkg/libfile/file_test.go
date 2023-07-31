package libFile

import (
	"fmt"
	"testing"
)

func Test_ReadFile(t *testing.T) {

	fmt.Println("读整个文件")
	fileName := "logs/test.log"
	data, err := ReadFile(fileName)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(string(data))
	}
}
func Test_ReadLineFile(t *testing.T) {

	fmt.Println("按行读整个文件")
	fileName := "logs/test.log"
	list, err := ReadLineFile(fileName)
	if err != nil {
		t.Error(err)
	} else {
		for _, v := range list {
			fmt.Println(string(v))
		}
	}
}

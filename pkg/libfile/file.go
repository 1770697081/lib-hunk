package libFile

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReadFile(filename string) ([]byte, error) {
	b, err := ioutil.ReadFile(filename)
	return b, err
}

// 按行读文件
func ReadLineFile(fileName string) ([]string, error) {

	listString := make([]string, 0, 50)
	if file, err := os.Open(fileName); err != nil {
		return listString, err
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			listString = append(listString, scanner.Text())
		}
	}
	return listString, nil
}

// 去掉空格和换行符
func ReplaceNilTab(str string) string {
	str = strings.Replace(str, " ", "", -1)  // 去除空格
	str = strings.Replace(str, "\n", "", -1) // 去除换行符
	str = strings.Replace(str, "\t", "", -1) // 去除换行符
	return str
}
func ToMd5(str string) string {
	return fmt.Sprintf("%02x", md5.Sum([]byte(str)))
}

// 将字符串转按行读取
func String2ByLine(data string) (array []string) {
	array = make([]string, 0, 100)
	reader := strings.NewReader(data)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {

		str := scanner.Text()
		str = strings.Replace(str, " ", "", -1)  // 去除空格
		str = strings.Replace(str, "\n", "", -1) // 去除换行符
		array = append(array, str)
	}
	return
}

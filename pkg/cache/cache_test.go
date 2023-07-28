package cache

import (
	"fmt"
	"testing"
)

func Test_PingProxy(t *testing.T) {

	fmt.Println("初始化本地缓存")

	fd := NewFreeCacheGlobal(10)
	fd.SetBytes("test", []byte("hahahah"), 10)

	v, e := fd.GetString("test")
	if e != nil {
		t.Error("测试失败:", e)
	} else {
		fmt.Println("测试成功:", v)
	}
}

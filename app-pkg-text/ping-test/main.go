package main

import (
	"github.com/1770697081/lib-hunk/pkg/logbeego"
	"github.com/astaxie/beego/logs"
)

func BeforeMainInit() {
	/**************************************************
	**1、所有初始化，尽量不要用init处理;都在此函数处理;
	***************************************************/
	logbeego.LogBeegoInit("cache-test", 10)
}
func main() {

	BeforeMainInit()

	logs.Info("嘻嘻哈哈")
	return
}

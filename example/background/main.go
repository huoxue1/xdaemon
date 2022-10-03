//本示例, 将把进程转为后台运行, 并保留所有参数不变

package main

import (
	"log"
	"os"
	"time"

	"github.com/zh-five/xdaemon"
)

func main() {
	//logFile := "daemon.log"

	//启动一个子进程后主程序退出
	xdaemon.Background(os.Stdout, true)

	//以下代码只有子程序会执行
	log.Println(os.Getpid(), "start...")
	time.Sleep(time.Second * 10)
	log.Println(os.Getpid(), "end")
}

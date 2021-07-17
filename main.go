package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
)

func main() {
	filePath := flag.String("f",
		"/Users/tinyhuiwang/workspaces/bjca.cn.ws/api-gateway-manager",
		"目录")

	realRemove := flag.Bool("rr",
		false,
		"是否真的要删除")
	flag.Parse()
	fmt.Println("----- file list start-----")
	files, _ := ioutil.ReadDir(*filePath)
	for _, f := range files {
		delTarget(f, *filePath, *realRemove)
	}
	fmt.Println("----- file list end-----")
}

func delTarget(current fs.FileInfo, parent string, rr bool) {
	PthSep := string(os.PathSeparator)
	currentAbs := parent + PthSep + current.Name()
	if current.Name() == "pom.xml" && !current.IsDir() {
		delTarget := parent + PthSep + "target"
		if Exists(delTarget) {
			fmt.Println("需要删除 target目录", delTarget)
			if rr {
				os.RemoveAll(delTarget)
				fmt.Println("删除 target目录", delTarget, "成功")
			}
		}
		return
	}
	if current.IsDir() && current.Name() != "target" {
		files, _ := ioutil.ReadDir(currentAbs)
		for _, f := range files {
			delTarget(f, currentAbs, rr)
		}
	}
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

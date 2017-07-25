package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	class       string
	args        []string
}

//解析命令行的工具方法
func parseCmd() *Cmd {
	cmd := &Cmd{}

}

//打印帮助信息
func printUsage() {
fmt.Printf("Usage: %s [-option] class [args]", ...)

}

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
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message.")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message.")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit.")
	flag.BoolVar(&cmd.versionFlag, "v", false, "print version and exit.")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.Parse() //解析失败就会调用printUsage

	args := flag.Args() //获取没有被解析的参数
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

//打印帮助信息
func printUsage() {
	fmt.Printf("Usage: %s [-option] class [args]\n", os.Args[0])
}

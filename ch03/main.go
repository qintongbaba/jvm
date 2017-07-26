package main

import (
	"fmt"
	"jvm/ch03/classpath"
	"strings"
)

func main() {
	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

//   jrePath : /Library/Java/JavaVirtualMachines/jdk1.8.0_122.jdk/Contents/Home/jre
func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("jreDir:%v classpath:%v class:%v args:%v\n", cmd.XjreOption, cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		fmt.Println(err)
		return
	}
	fmt.Printf("class data:%v\n", classData)
}

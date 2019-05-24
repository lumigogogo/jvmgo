package main

import (
	"fmt"
	"jvmgo/ch02/classpath"
	"strings"
)

func main()  {
	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n", cp, cmd.class, cmd.args)
	classname := strings.Replace(cmd.class, ".", "/", -1)
	fmt.Printf("classname:%s \n", classname)
	classData, _, err := cp.ReadClass(classname)
	if err != nil {
		fmt.Printf("class not found or load main class %s\n", cmd.class)
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("class data: %v \n", classData)
}

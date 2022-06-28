package main

import (
	"fmt"
	"jvmgo/ch05/classfile"
	"jvmgo/ch05/classpath"
	"strings"
)

/*/Users/xiuwenyin/go/bin/ch05 -Xjre "/Library/Java/JavaVirtualMachines/jdk1.8.0_271.jdk/Contents/Home/jre/" java.lang.String*/

func main() {
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
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range cf.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}

////testLocalVars()函数测试局部变量
//func testLocalVars(vars *rtda.LocalVars) {
//	vars.SetInt(0, 100)
//	vars.SetInt(1, -100)
//	vars.SetLong(2, 2997924580)
//	vars.SetLong(4, -2997924580)
//	vars.SetFloat(6, 3.1415926)
//	vars.SetDouble(7, 2.71828182845)
//	vars.SetRef(9, nil)
//	println(vars.GetInt(0))
//	println(vars.GetInt(1))
//	println(vars.GetLong(2))
//	println(vars.GetLong(4))
//}
//
////testOperandStack()函数测试操作数栈
//func testOperandStack(ops *rtda.OperandStack) {
//	ops.PushInt(100)
//	ops.PushInt(-100)
//	ops.PushLong(2997924580)
//	ops.PushLong(-2997924580)
//	ops.PushFloat(3.1415926)
//	ops.PushDouble(2.71828182845)
//	ops.PushRef(nil)
//	println(ops.PopRef())
//	println(ops.PopDouble())
//	println(ops.PopFloat())
//	println(ops.PopLong())
//	println(ops.PopLong())
//	println(ops.PopInt())
//	println(ops.PopInt())
//}
//
//func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
//	classData, _, err := cp.ReadClass(className)
//	if err != nil {
//		panic(err)
//	}
//	cf, err := classfile.Parse(classData)
//	if err != nil {
//		panic(err)
//	}
//	return cf
//}
//
//func printClassInfo(cf *classfile.ClassFile) {
//	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
//	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
//	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
//	fmt.Printf("this class: %v\n", cf.ClassName())
//	fmt.Printf("super class: %v\n", cf.SuperClassName())
//	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
//	fmt.Printf("fields count: %v\n", len(cf.Fields()))
//	for _, f := range cf.Fields() {
//		fmt.Printf(" %s\n", f.Name())
//	}
//	fmt.Printf("methods count: %v\n", len(cf.Methods()))
//	for _, m := range cf.Methods() {
//		fmt.Printf(" %s\n", m.Name())
//	}
//}

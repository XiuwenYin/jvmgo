package base

import "jvmgo/ch05/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader) //FetchOperands()方法从字节码中提取操作数
	Execute(frame *rtda.Frame)            //Execute()方法 执行指令逻辑
}

//NoOperandsInstruction 表示没有操作数的指令，所以没有定义任何字段
type NoOperandsInstruction struct{}

//BranchInstruction 表示跳转指令
type BranchInstruction struct {
	Offset int //Offset 字段存放跳转偏移量
}

// Index8Instruction 存储和加载类指令需要根据索引存取局部变量表，索引由单字节操作数给出。把这类指令抽象成Index8Instruction结构体
type Index8Instruction struct {
	Index uint //Index 字段表示局部变量表索引
}

// Index16Instruction 有一些指令需要访问运行时常量池，常量池索引由两字节操作数给出。把这类指令抽象成Index16Instruction结构体
type Index16Instruction struct {
	Index uint
}

//FetchOperands 方法自然也是空空如也，什么也不用读
func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

//FetchOperands ()方法从字节码中读取一个uint16整数，转成int后赋 给Offset字段
func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

//FetchOperands 方法从字节码中 读取一个int8整数，转成uint后赋给Index字段
func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

//FetchOperands 方法从字节码中读取一个 uint16整数，转成uint后赋给Index字段
func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}

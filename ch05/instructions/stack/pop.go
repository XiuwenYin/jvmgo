package stack

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

/*栈指令直接对操作数栈进行操作, pop和pop2指令将栈顶变量弹出*/

// POP 指令只能用于弹出int、float等占用一个操作数栈位置的变量
type POP struct{ base.NoOperandsInstruction }

// POP2 double和long变量在操作数栈中占据两个位置，需要使用POP2指令弹出
type POP2 struct{ base.NoOperandsInstruction }

func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}

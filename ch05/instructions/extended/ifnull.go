package extended

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// 根据引用是否是null进行跳转，ifnull和ifnonnull指令把栈顶的引用弹出

type IFNULL struct{ base.BranchInstruction } // Branch if reference is null

func (self *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}

type IFNONNULL struct{ base.BranchInstruction } // Branch if reference not null

func (self *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}

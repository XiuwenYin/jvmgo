package constants

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type NOP struct {
	base.BranchInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	// 什么也不用做
}

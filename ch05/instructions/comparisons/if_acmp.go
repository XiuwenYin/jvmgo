package comparisons

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

// Branch if reference comparison succeeds

// IF_ACMPEQ 和 IF_ACMPNE 指令把栈顶的两个引用弹出，根据引用是否相同进行跳转
type IF_ACMPEQ struct{ base.BranchInstruction }

func (self *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	if _acmp(frame) {
		base.Branch(frame, self.Offset)
	}
}

type IF_ACMPNE struct{ base.BranchInstruction }

func (self *IF_ACMPNE) Execute(frame *rtda.Frame) {
	if !_acmp(frame) {
		base.Branch(frame, self.Offset)
	}
}

func _acmp(frame *rtda.Frame) bool {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	return ref1 == ref2
}

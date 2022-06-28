package rtda

type Frame struct {
	lower        *Frame        //lower字段用来实现链表数据结构
	localVars    LocalVars     //localVars字段保存 局部变量表指针
	operandStack *OperandStack //operandStack字段保存操作数栈指针
	thread       *Thread
	nextPC       int
}

//NewFrame 函数创建Frame实例
func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

// getters & setters

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
func (self *Frame) Thread() *Thread {
	return self.thread
}
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}

package rtda

type Frame struct {
	lower        *Frame        //lower字段用来实现链表数据结构
	localVars    LocalVars     //localVars字段保存 局部变量表指针
	operandStack *OperandStack //operandStack字段保存操作数栈指针

}

//NewFrame 函数创建Frame实例
func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack)}
}

func (this *Frame) Lower() *Frame {
	return this.lower
}

func (this *Frame) OperandStack() *OperandStack {
	return this.operandStack
}

func (this *Frame) LocalVars() *LocalVars {
	return &this.localVars
}

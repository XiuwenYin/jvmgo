package rtda

/*Thread 目前只定义了pc和stack两个字段。
pc字段无需解释，stack字段是Stack结构体(Java虚拟机栈)指针,
java命令提供了-Xss选项来设置Java虚拟机栈大小*/
/*
JVM
  Thread
    pc
    Stack
      Frame
        LocalVars
        OperandStack
*/
type Thread struct {
	pc    int // the address of the instruction currently being executed
	stack *Stack
}

//NewThread 函数创建Thread实例
func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

// PC getter方法
func (self *Thread) PC() int {
	return self.pc
}

// SetPC setter方法
func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

// PushFrame 调用Stack结构体的相应方法
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

//PopFrame 调用Stack结构体的相应方法
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

//CurrentFrame 方法返回当前帧
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) NewFrame(maxLocals, maxStack uint) *Frame {
	return newFrame(self, maxLocals, maxStack)
}

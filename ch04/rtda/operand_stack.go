package rtda

import "math"

// OperandStack 操作数栈的大小是编译器已经确定的，所以可以用[]Slot实现
type OperandStack struct {
	size  uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

/*和局部变量表类似，需要定义一些方法从操作数栈中弹出，或者往其中推入各种类型的变量*/

//PushInt 方法往栈顶放一个int变量，然后把size加1
func (self *OperandStack) PushInt(val int32) {
	self.slots[self.size].num = val
	self.size++
}

//PopInt 方法则恰好相反，先把size减1，然后返回变量值
func (self *OperandStack) PopInt() int32 {
	self.size--
	return self.slots[self.size].num
}

// PushFloat float变量还是先转成int类型，然后按int变量处理
func (self *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	self.slots[self.size].num = int32(bits)
	self.size++
}
func (self *OperandStack) PopFloat() float32 {
	self.size--
	bits := uint32(self.slots[self.size].num)
	return math.Float32frombits(bits)
}

// PushLong 把long变量推入栈顶时，要拆成两个int变量
func (self *OperandStack) PushLong(val int64) {
	self.slots[self.size].num = int32(val)
	self.slots[self.size+1].num = int32(val >> 32)
	self.size += 2
}

// PopLong 弹出时，先弹出两个int变量，然后组装成一个long变量
func (self *OperandStack) PopLong() int64 {
	self.size -= 2
	low := uint32(self.slots[self.size].num)
	high := uint32(self.slots[self.size+1].num)
	return int64(high)<<32 | int64(low)
}

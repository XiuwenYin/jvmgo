package control

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

/*
tableswitch
<0-3 byte pad>
defaultbyte1
defaultbyte2
defaultbyte3
defaultbyte4
lowbyte1
lowbyte2
lowbyte3
lowbyte4
highbyte1
highbyte2
highbyte3
highbyte4
jump offsets...
*/
// Access jump table by index and jump
type TABLE_SWITCH struct {
	defaultOffset int32 // defaultOffset 对应默认情况下执行跳转所需的字节码偏移量
	low           int32 // low 和 high 记录case的取值范围
	high          int32
	jumpOffsets   []int32 // jumpOffsets 是一个索引表，里面存放 high - low + 1 个int值，对应各种case情况下，执行跳转所需的字节码偏移量
}

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

/*Execute 方法先从操作数栈中弹出一个int变量，然后看它是 否在low和high给定的范围之内。
如果在，则从jumpOffsets表中查出 偏移量进行跳转，否则按照defaultOffset跳转
*/
func (self *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffsets[index-self.low])
	} else {
		offset = int(self.defaultOffset)
	}
	base.Branch(frame, offset)
}

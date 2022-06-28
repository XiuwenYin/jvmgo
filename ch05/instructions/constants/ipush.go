package constants

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type BIPUSH struct{ val int8 } // Push byte

/*bipush指令从操作数中获取一个byte型整数，扩展成int型，然后推入栈顶*/
func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(int32(self.val))
}

type SIPUSH struct{ val int16 } // Push short

/*sipush指令从操作数中获取一个short型整数，扩展成int型，然后推入栈顶*/
func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}

func (self *SIPUSH) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(int32(self.val))
}

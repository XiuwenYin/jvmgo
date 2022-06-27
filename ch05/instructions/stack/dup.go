package stack

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

/*栈指令直接对操作数栈进行操作, dup系列指令复制栈顶变量*/

type DUP struct{ base.NoOperandsInstruction }
type DUP_X1 struct{ base.NoOperandsInstruction }
type DUP_X2 struct{ base.NoOperandsInstruction }
type DUP2 struct{ base.NoOperandsInstruction }
type DUP2_X1 struct{ base.NoOperandsInstruction }
type DUP2_X2 struct{ base.NoOperandsInstruction }

// Execute DUP复制栈顶的单个变量
func (self *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

// Execute 将原本的 val1 >> val2 >> ... 变为 val1 >> val2 >> val1 >> ... 不可用与val2为long和double的情况
func (self *DUP_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopSlot()
	val2 := stack.PopSlot()
	stack.PushSlot(val1)
	stack.PushSlot(val2)
	stack.PushSlot(val1)
}

// Execute 将原本的 val1 >> val2 >> val3 >> ... 变为 val1 >> val2 >> val3 >> val1 >> ...
func (self *DUP_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// Execute 对于long和double，由于占据操作数栈的两个位置，所以pop两个slots，再各push两个回去
func (self *DUP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}

func (self *DUP2_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

/*Execute slot1和slot2应为1个操作数，slot3和slot4可为1个或2个操作数，
val1 >> val2 >> val3 >> val4 >> ... ->
val2 >> val1 >> val4 >> val3 >> val2 >> val1 >> ...*/
func (self *DUP2_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	slot4 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot4)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)

}

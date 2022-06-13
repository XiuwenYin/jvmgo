package base

type BytecodeReader struct {
	code []byte //code 字段存放字节码
	pc   int    //pc 字段记录读取到了哪个字节
}

// Reset 为了避免每次解码指令都新创建一个BytecodeReader实例，给它定义一个 Reset()方法
func (self *BytecodeReader) Reset(code []byte, pc int) {
	self.code = code
	self.pc = pc
}

func (self *BytecodeReader) ReadUint8() uint8 {
	i := self.code[self.pc]
	self.pc++
	return i
}

//ReadInt8 方法调用ReadUint8()，然后把读取到的值转成int8 返回
func (self *BytecodeReader) ReadInt8() int8 {
	return int8(self.ReadUint8())
}

//ReadUint16 连续读取两字节
func (self *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(self.ReadUint8())
	byte2 := uint16(self.ReadUint8())
	return (byte1 << 8) | byte2
}

//ReadInt16 方法调用ReadUint16()，然后把读取到的值转成 int16返回
func (self *BytecodeReader) ReadInt16() int16 {
	return int16(self.ReadUint16())
}

func (self *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(self.ReadUint8())
	byte2 := int32(self.ReadUint8())
	byte3 := int32(self.ReadUint8())
	byte4 := int32(self.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}
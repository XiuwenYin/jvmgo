package classfile

import "math"

type ConstantIntegerInfo struct {
	val int32
}

type ConstantFloatInfo struct {
	val float32
}

type ConstantLongInfo struct {
	val int64
}

type ConstantDoubleInfo struct {
	val float64
}

// readInfo()先读取一个uint32数据，然后把它转型成int32类型
func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

// readInfo()先读取一个uint32数据，然后调用math包的 Float32frombits()函数把它转换成float32类型
func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes)
}

// readInfo()先读取一个uint64数据，然后把它转型成int64类型
func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}

// readInfo()先读取一个uint64数据，然后调用math包的 Float64frombits()函数把它转换成float64类型
func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = math.Float64frombits(bytes)
}

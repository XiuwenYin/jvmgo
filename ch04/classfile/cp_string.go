package classfile

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

// readInfo()方法读取常量池索引
func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}

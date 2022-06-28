package classfile

import "fmt"

/*常量池实际上也是一个表，但是有三点需要特别注意。
第一，表头给出的常量池大小比实际大1。假设表头给出的值是n，那么常 量池的实际大小是n–1。
第二，有效的常量池索引是1~n–1。0是无效 索引，表示不指向任何常量。
第三，CONSTANT_Long_info和 CONSTANT_Double_info各占两个位置。
也就是说，如果常量池中存在这两种常量，实际的常量数量比n–1还要少，而且1~n–1的某些 数也会变成无效索引*/

type ConstantPool []ConstantInfo

//type ConstantInfo interface {
//	readInfo(reader *ClassReader)
//}

// readConstantPool 读取常量池
func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ { // 注意索引从1开始
		cp[i] = readConstantInfo(reader, cp)
		// http://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4.5
		// All 8-byte constants take up two entries in the constant_pool table of the class file.
		// If a CONSTANT_Long_info or CONSTANT_Double_info structure is the item in the constant_pool
		// table at index n, then the next usable item in the pool is located at index n+2.
		// The constant_pool index n+1 must be valid but is considered unusable.
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++ // 占两个位置
		}
	}
	return cp
}

// getConstantInfo 方法按索引查找常量
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic(fmt.Errorf("invalid constant pool index: %v", index))
}

// getNameAndType 方法从常量池查找字段或方法的名字和描述符
func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

// getClassName 方法从常量池查找类名
func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

// getUtf8 方法从常量池查找UTF-8字符串
func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}

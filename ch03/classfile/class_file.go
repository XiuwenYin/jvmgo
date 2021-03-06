package classfile

import "fmt"

// ClassFile
/*尽量只公开必要的变量、字段、函数和 方法等。但是为了提高代码可读性，所有的结构体都是公开的*/
type ClassFile struct {
	//magic uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

// Parse 函数把[]byte解析成ClassFile结构体
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

// Go语言没有异常处理机制，只有一个panic-recover机制。
// read 方法依次调用其他方法解析class文件
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

// class文件的魔数 是"0xCAFEBABE"
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		// 刚刚开始编写虚拟机，还无法抛出异常，所以暂时先调用 panic()方法终止程序执行
		panic("java.lang.ClassFormatError: magic!")
	}
}

/*readAndCheckVersion
class文件的次版本号和主版本号，都是u2类型。
假设某class文件的主版本号是M，次版本号是m，那么完整的版本号可以表示成 "M.m" 的形式*/
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

// MinorVersion getter方法
func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

//MajorVersion 等6个方法是Getter方法，把结构体的字段暴露 给其他包使用
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

// ConstantPool  getter方法
func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}

// AccessFlags getter方法
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

// Fields getter方法
func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}

// Methods getter方法
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

// ClassName 从常量池查找类名
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

// SuperClassName 从常量池查找超类名
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return "" // 只有 java.lang.Object 没有超类

}

// InterfaceNames 常量池查找接口名
func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

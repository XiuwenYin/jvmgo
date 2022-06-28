package classfile

var (
	_attrDeprecated = &DeprecatedAttribute{}
	_attrSynthetic  = &SyntheticAttribute{}
)

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/

/*属性表中存放的属性名实际上并不是编码后的字符串， 而是常量池索引，指向常量池中的CONSTANT_Utf8_info常量*/

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

// readAttributes()函数读取属性表
func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

// readAttribute()读取单个属性
/*先读取属性名索引，根据它从常量池中找到属性名，
然后读取属性长度，接着调用newAttributeInfo()函数创建具体的属性实例*/
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "BootstrapMethods":
		return &BootstrapMethodsAttribute{}
	case "Code": //方法里的Java代码，经过Javac编译器编译成字节码指令之后，存放在方法属性表集合中一个名为“Code”的属性里面
		return &CodeAttribute{cp: cp}
	case "ConstantValue": //由final关键字定义的常量值
		return &ConstantValueAttribute{}
	case "Deprecated": //被声明为 deprecated 的方法与字段
		return &DeprecatedAttribute{}
	case "EnclosingMethod":
		return &EnclosingMethodAttribute{cp: cp}
	case "Exceptions": //方法抛出的异常列表
		return &ExceptionsAttribute{}
	case "InnerClasses":
		return &InnerClassesAttribute{}
	case "LineNumberTable": //Java原码的行号与字节码指令的对应关系
		return &LineNumberTableAttribute{}
	case "LocalVariableTable": //方法局部变量的描述
		return &LocalVariableTableAttribute{}
	case "LocalVariableTypeTable":
		return &LocalVariableTypeTableAttribute{}
	//case "MethodParameters":
	//case "RuntimeInvisibleAnnotations":
	//case "RuntimeInvisibleParameterAnnotations":
	//case "RuntimeInvisibleTypeAnnotations":
	//case "RuntimeVisibleAnnotations":
	//case "RuntimeVisibleParameterAnnotations":
	//case "RuntimeVisibleTypeAnnotations":
	case "Signature":
		return &SignatureAttribute{cp: cp}
	case "SourceFile": //记录源文件名称
		return &SourceFileAttribute{cp: cp}
	case "Synthetic": //标识方法或字段为编译器自动生成的
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}

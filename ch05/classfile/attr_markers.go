package classfile

/*
Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/

type DeprecatedAttribute struct {
	MarkerAttribute
}

/*
Synthetic_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/

type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct{}

//由于这两个属性都没有数据，所以 readInfo()方法是空的
func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}

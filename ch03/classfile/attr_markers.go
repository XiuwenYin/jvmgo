package classfile

type DeprecatedAttribute struct{ MarkerAttribute }
type SyntheticAttribute struct{ MarkerAttribute }
type MarkerAttribute struct{}

//由于这两个属性都没有数据，所以 readInfo()方法是空的
func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}

package rtda

import "math"

/*局部变量表*/
/*并没有对boolean、byte、short和char类型定义存取方法，这些类型的值都可以转换成int值类来处理*/

type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

//SetInt set方法
func (self LocalVars) SetInt(index uint, val int32) {
	self[index].num = val
}

//GetInt get方法
func (self LocalVars) GetInt(index uint) int32 {
	return self[index].num
}

// SetFloat float变量可以先转成int类型，然后按int变量来处理。
func (self LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self[index].num = int32(bits)
}
func (self LocalVars) GetFloat(index uint) float32 {
	bits := uint32(self[index].num)
	return math.Float32frombits(bits)
}

// SetLong long变量则需要拆成两个int变量
func (self LocalVars) SetLong(index uint, val int64) {
	self[index].num = int32(val)
	self[index+1].num = int32(val >> 32)
}
func (self LocalVars) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index+1].num)
	return int64(high)<<32 | int64(low)
}

// SetDouble double变量可以先转成long类型，然后按照long变量来处理
func (self LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}
func (self LocalVars) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}

// SetRef 最后是引用值，直接存取即可
func (self LocalVars) SetRef(index uint, ref *Object) {
	self[index].ref = ref
}
func (self LocalVars) GetRef(index uint) *Object {
	return self[index].ref
}

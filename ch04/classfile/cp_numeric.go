package classfile

import (
	"math"
)

//CONSTANT_Integer_info正好可以容纳一个Java的int型常量，
//但实际上比int更小的boolean、byte、short和char类型的常量也放在CONSTANT_Integer_info中
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

func (self *ConstantIntegerInfo) readInfo(read *ClassReader) {
	bytes := read.readUint32()
	self.val = int32(bytes)
}

func (self *ConstantFloatInfo) readInfo(read *ClassReader) {
	bytes := read.readUint32()
	self.val = math.Float32frombits(bytes)
}

func (self *ConstantLongInfo) readInfo(read *ClassReader) {
	bytes := read.readUint64()
	self.val = int64(bytes)
}

func (self *ConstantDoubleInfo) readInfo(read *ClassReader) {
	bytes := read.readUint64()
	self.val = math.Float64frombits(bytes)
}

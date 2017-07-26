package classfile

type ConstantPool []ConstantInfo

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantPool(reader *ClassReader) ConstantPool {}

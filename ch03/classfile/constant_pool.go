package classfile

type ConstantPool []ConstantInfo

//常量池信息接口    因为常量池信息分为多种
type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

//读取常量池
/*
第一，表头给出的常量池大小比实际大1。假设表头给出的值是n，那么常量池的实际大小是n–1。
第二，有效的常量池索引是1~n–1。0是无效索引，表示不指向任何常量。
第三，CONSTANT_Long_info和CONSTANT_Double_info各占两个位置。也就是说，如果常量池中
      存在这两种常量，实际的常量数量比n–1还要少，而且1~n–1的某些数也会变成无效索引
*/
func readConstantPool(reader *ClassReader) ConstantPool {
	//获取常量池大小
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ {

	}
}

func readContantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {

}

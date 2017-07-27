package classfile

import (
	"fmt"
)

type ClassFile struct {
	//magic   uint32
	minorVersion uint16       //次版本号
	majorVersion uint16       //主版本号
	constantPool ConstantPool //常量池
	accessFlags  uint16       //类访问标志
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return cf, nil
}
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)               //读取魔数
	self.readAndCheckVersion(reader)             //检查其支持版本
	self.constantPool = readConstantPool(reader) //读取常量池
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()

	//检查是否为class的魔数
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}
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
func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
} //getter

func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
} //getter

func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
} //getter

func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
} //getter

func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
} //getter

func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
} //getter

func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return "" //只有java.lang.Object没有超类
}
func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

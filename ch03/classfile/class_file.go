package classfile

import (
	"fmt"
)

type ClassFile struct {
	//magic   uint32

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

func Parse(classData []byte) (cf *ClassFile, err error)         {}
func (self *ClassFile) read(reader *ClassReader)                {}
func (self *ClassFile) readAndCheckMagic(reader *ClassReader)   {}
func (self *ClassFile) readAndCheckVwesion(reader *ClassReader) {}
func (self *ClassFile) MinorVersion() uint16                    {} //getter
func (self *ClassFile) MajorVersion() uint16                    {} //getter
func (self *ClassFile) ConstantPool() ConstantPool              {} //getter
func (self *ClassFile) AccessFlags() uint16                     {} //getter
func (self *ClassFile) Fields() []*MemberInfo                   {} //getter
func (self *ClassFile) Methods() []*MemberInfo                  {} //getter
func (self *ClassFile) ClassName() string                       {}
func (self *ClassFile) SuperClassName() string                  {}
func (self *ClassFile) InterfaceNames() []string                {}

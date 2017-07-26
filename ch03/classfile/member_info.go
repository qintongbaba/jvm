package classfile

type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readmember(reader, cp)
	}
	return members
}
func readmember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}
func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}
func (self *MemberInfo) Names() string {

}
func (self *MemberInfo) Descriptor() string {

}

package classfile

// 字段和方法
type MemberInfo struct {
	cp              ConstantPool // cp 字段保存常量池指针
	accessFlags     uint16       // 访问标志
	nameIndex       uint16       //
	descriptorIndex uint16
	attributes      []AttributeInfo
}

// readMembers() 读取字段表或方法表
func readMembers(read *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

// readMember（）函数读取字段或方法数据
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:                         cp,
		accessFlags:                reader.readUint16(),
		nameIndex:                  reader.readUint16(),
		descriptorIndex:            reader.readUint16(),
		readAttributes(reader, cp), // 见3.4属性表和readAttributes()函数
	}
}
func (self *MemberInfo) AccessFlags() uint16 { //getter
	return self.accessFlags
}

// 从常量池查找字段或方法名
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

// 从常量池查找字段或方法描述符
func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}

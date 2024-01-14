package classfile

import "fmt"

// Go的访问控制非常简单：只有公开和私有两种。
// 所有首字母大写的类型、结构体、字段、变量、函数、方法都是公开的，可供其他包使用。
// 首字母小写则是私有的，只能在包内部使用
// Go语言没有异常处理机制，只有一个panic-recover机制
type ClassFile struct { // 反映Java虚拟机规范定义的class文件格式。

	// magic uint32
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

// Parse函数 把[]byte 解析成ClassFile结构体
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

// read方法依次调用其他方法解析class文件
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)               // 见3.2.3
	self.readAndCheckVersion(reader)             // 见3.2.4
	self.constantPool = readConstantPool(reader) // 3.3
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.intefaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool) // 见3.2.8
	self.methods = readMembers(self, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError:magic!")
	}
}

// 支持版本号为45.0~52.0的class文件。如果其他版本号，暂时先调用panic()方法终止程序执行
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.minorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

// MajorVersion方法 等六个方法是Getter 方法，把结构体的字段暴露给其他包使用
func (self *ClassFile) MinorVersion() uint16 { //getter
	return self.minorVersion
}
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}
func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}
func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

// ClassName从常量池查找类名
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return "" // 只有java.lang.Object没有超类
}
func (self *ClassFile) InterfaceNames() []string {

	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

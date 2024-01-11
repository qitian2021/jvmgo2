package classfile

import (
	"encoding/binary"
)

type ClassReader struct {
	data []byte
}

// ClassReader并没有使用索引记录数据位置，而是使用Go语言的reslice语法跳过已经读取的数据
func (self *ClassReader) readUint8() uint8 { // u1
	val := self.data[0]
	self.data = self.data[1:]
	return val
}
func (self *ClassReader) readUint16() uint16 { // u2
	val := binary.BigEndian.Uint16(self.data) // BigEndian,可以从[]byte 中解码多字节数据。
	self.data = self.data[2:]
	return val
}

func (self *ClassReader) readUint32() uint32 { // u4
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}
func (self *ClassReader) readUint64() uint64 { // u8 Java虚拟机规范并没有定义u8类型数据
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}
func (self *ClassReader) readUint16s() []uint16 { // 读取uint16表，表的大小由开头的uint16数据指出
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}
func (self *ClassReader) readBytes(length uint32) []byte { // 读取指定数量的字节
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes

}

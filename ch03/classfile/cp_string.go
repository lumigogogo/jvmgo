package classfile

/*
	CONSTANT_String_info {
	   u1 tag;
	   u2 string_index;
	}

CONSTANT_String_info本身并不存放字符串数据，只存了常量池索引，这个索引指向一个CONSTANT_Utf8_info常量
*/
type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

func (self *ConstantStringInfo) string() string {
	return self.cp.getUtf8(self.stringIndex)
}

package classfile

/*
	CONSTANT_Utf8_info {
	   u1 tag;
	   u2 length;
	   u2 length;
	   u1 bytes[length];
	}

字符串在class文件中是以MUTF-8（Modified UTF-8）方式编码的
*/
type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}

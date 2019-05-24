package classfile

type UnparsedAttribute struct {
	name string
	len  uint32
	info []byte
}

func (self *UnparsedAttribute) readInfo(reader *ClassReader) {
	self.info = reader.readBytes(self.len)
}

package classfile

/* 属性表接口 */
/*
	attribute_info {
	   u2 attribute_name_index;
	   u4 attribute_length;
	   u1 info[attribute_length];
	}
*/
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

/* 读取属性表 */
func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := NewAttribute(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

func NewAttribute(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	/*
		23种预定义属性分为3组，第一组是实现虚拟机所必须得，共5种，
		第二组是java类库所必须得，共有12种；
		第三组属性提供给工具使用共6种，是可选的
	 */
	switch attrName {
	case "Code":
		return &CodeAttribute{cp: cp}	// 1组
	case "ConstantValue":
		return &ConstantValueAttribute{}	// 1组
	case "Deprecated":
		return &DeprecatedAttribute{}	// 3组
	case "Exceptions":
		return &ExceptionsAttribute{}	// 1组
	case "LineNumberTable":
		return &LineNumberTableAttribute{}	//3组 在异常堆栈中显示行号
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}	// 3组
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}	// 3组
	case "Synthetic":
		return &SyntheticAttribute{}	// 2组
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}

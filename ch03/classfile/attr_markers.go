package classfile

/*
Deprecated和Synthetic是最简单的两种属性，仅起标记作用，不包含任何数据。
这两种属性都是JDK1.1引入的，可以出现在ClassFile、field_info和method_info结构中
*/

/*
	Deprecated_attribute {
	   u2 attribute_name_index;
	   u4 attribute_length;
	}
	Synthetic_attribute {
	   u2 attribute_name_index;
	   u4 attribute_length;
	}
*/
type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct {
}

/* 由于这两个属性都没有数据，所以readInfo（）方法是空的 */
func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}

package rtda

type Frame struct {
	lower        *Frame        // 用来实现链表数据结构
	LocalVars    LocalVars     // 保存局部变量表指针
	OperandStank *OperandStack // 保存操作数栈指针
}

/*
执行方法所需的局部变量表大小和操作数栈深度是由编译器预先计算好的，
存储在class文件method_info结构的Code属性中，具体可以参考3.4.5节。
*/
func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		LocalVars:    newLocalVars(maxLocals),
		OperandStank: newOperandStack(maxStack),
	}
}

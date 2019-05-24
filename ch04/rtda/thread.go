package rtda

type Thread struct {
	pc    int    // pc寄存器
	stack *Stack // JVM栈
}

func NewTrhead() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

// getter
func (self *Thread) PC() int {
	return self.pc
}

// setter
func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

/* 入栈 */
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

/* 出栈 */
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

/* 当前帧 */
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}

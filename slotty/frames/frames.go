package frames

import (
	"bakeoff/slotty/object"
	"container/list"
)

type Frame struct {
	Thread       int
	FrameStack   *list.List  // points to the frame stack
	MethName     string      // method name
	MethType     string      // method type (signature)
	ClName       string      // class name
	Meth         []byte      // bytecode of method
	CP           interface{} // will hold a *classloader.CPool (constant pool ptr) but due to circularity must be done this way
	Locals       []Slot      // local variables
	OpStack      []Slot      // operand stack
	TOS          int         // top of the operand stack
	PC           int         // program counter (index into the bytecode of the method)
	Ftype        byte        // type of method in frame: 'J' = java, 'G' = Golang, 'N' = native
	ExceptionPC  int         // program counter at the moment the PC threw an exception
	WideInEffect bool        // WideInEffect indicates if the wide instruction is in effect in the current frame
	AccessFlags  int         // access flags for the method
	// ObjSync is a reference to the object whose method is being executed in this frame.
	// It is needed for synchronization purposes.
	ObjSync *object.Object // Object for method synchronization
}

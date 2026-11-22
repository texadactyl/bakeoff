package frames

type Frame struct {
	OpStack []any
	Locals  []any
	TOS     int // index of top of stack
}

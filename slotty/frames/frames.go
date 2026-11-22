package frames

import "bakeoff/slotty/pslot"

type Frame struct {
	OpStack []pslot.Slot
	Locals  []pslot.Slot
	TOS     int // index of top of stack
}

package slotty

import (
	"bakeoff/slotty/frames"
	"bakeoff/slotty/object"
	"bakeoff/slotty/pslot"
	"math"
)

func push(fr *frames.Frame, slot pslot.Slot) {
	fr.TOS++
	fr.OpStack[fr.TOS] = slot
}

func pop(fr *frames.Frame) pslot.Slot {
	slot := fr.OpStack[fr.TOS]
	fr.TOS--
	return slot
}

func peek(fr *frames.Frame) pslot.Slot {
	return fr.OpStack[fr.TOS]
}

func doIadd(fr *frames.Frame, _ int64) int {
	v2 := pslot.AsInt32(pop(fr))
	v1 := pslot.AsInt32(pop(fr))
	push(fr, pslot.IntSlot(int64(v1+v2)))
	return 1
}

func doLadd(fr *frames.Frame, _ int64) int {
	v2 := pslot.AsInt64(pop(fr))
	v1 := pslot.AsInt64(pop(fr))
	push(fr, pslot.IntSlot(v1+v2))
	return 1
}

func doFadd(fr *frames.Frame, _ int64) int {
	rhs := pslot.AsFloat32(pop(fr))
	lhs := pslot.AsFloat32(pop(fr))
	push(fr, pslot.Float32Slot(lhs+rhs))
	return 1
}

func doDadd(fr *frames.Frame, _ int64) int {
	rhs := pslot.AsFloat64(pop(fr))
	lhs := pslot.AsFloat64(pop(fr))
	push(fr, pslot.Float64Slot(lhs+rhs))
	return 1
}

func doDup(fr *frames.Frame, _ int64) int {
	tosItem := peek(fr)
	push(fr, tosItem)
	return 1
}

func doDup2(fr *frames.Frame, _ int64) int {
	top := pop(fr)
	next := peek(fr)
	push(fr, top)
	push(fr, next)
	push(fr, top)
	return 1
}

func Slotty(numLoops int) bool {
	var slot pslot.Slot
	var emptyOpStack []pslot.Slot
	var fr frames.Frame
	var pi64 = 3.14
	var pi64p1 = 4.14
	var one32 float32 = 1.0
	var one64 = 1.0
	var pi32 = float32(3.14)

	for range 10 {
		emptyOpStack = append(emptyOpStack, pslot.IntSlot(0))
	}
	var locals []pslot.Slot

	for ix := range numLoops {

		fr = frames.Frame{OpStack: emptyOpStack, Locals: locals, TOS: -1}

		// doIadd
		push(&fr, pslot.IntSlot(math.MaxInt32-1))
		push(&fr, pslot.IntSlot(1))
		doIadd(&fr, 0)
		slot = pop(&fr)

		// doLadd
		push(&fr, pslot.IntSlot(math.MaxInt64-1))
		push(&fr, pslot.IntSlot(1))
		doLadd(&fr, 0)
		slot = pop(&fr)

		// doFadd
		push(&fr, pslot.Float32Slot(pi32))
		push(&fr, pslot.Float32Slot(one32))
		doFadd(&fr, 0)
		slot = pop(&fr)

		// doDadd
		push(&fr, pslot.Float64Slot(pi64))
		push(&fr, pslot.Float64Slot(one64))
		doDadd(&fr, 0)
		slot = pop(&fr)

		if ix&0x1 == 0 {
			// doDup
			push(&fr, pslot.Float64Slot(pi64))
			doDup(&fr, 0)
			slot = pop(&fr)
			slot = pop(&fr)

			// doDup2
			push(&fr, pslot.Float64Slot(pi64))
			push(&fr, pslot.Float64Slot(pi64p1))
			doDup2(&fr, 0)
			slot = pop(&fr)
			slot = pop(&fr)
			slot = pop(&fr)
			slot = pop(&fr)
		}

		// Object
		obj := &object.Object{X: 1, Y: 2, Z: 3}
		push(&fr, pslot.ObjSlot(obj))
		slot = pop(&fr)
		obj = pslot.AsObject(slot)

	} // for ix := range numLoops

	// Success!
	return true
}

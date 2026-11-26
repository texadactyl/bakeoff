package slotty

import (
	"bakeoff/slotty/frames"
	"bakeoff/slotty/object"
	"math"
)

func gfunc(params []any) any {
	return params[0]
}

func push(fr *frames.Frame, slot frames.Slot) {
	fr.TOS++
	fr.OpStack[fr.TOS] = slot
}

func pop(fr *frames.Frame) frames.Slot {
	slot := fr.OpStack[fr.TOS]
	fr.TOS--
	return slot
}

func peek(fr *frames.Frame) frames.Slot {
	return fr.OpStack[fr.TOS]
}

func doIadd(fr *frames.Frame, _ int64) int {
	v2 := frames.AsInt32(pop(fr))
	v1 := frames.AsInt32(pop(fr))
	push(fr, frames.IntSlot(int64(v1+v2)))
	return 1
}

func doLadd(fr *frames.Frame, _ int64) int {
	v2 := frames.AsInt64(pop(fr))
	v1 := frames.AsInt64(pop(fr))
	push(fr, frames.IntSlot(v1+v2))
	return 1
}

func doFadd(fr *frames.Frame, _ int64) int {
	rhs := frames.AsFloat32(pop(fr))
	lhs := frames.AsFloat32(pop(fr))
	push(fr, frames.Float32Slot(lhs+rhs))
	return 1
}

func doDadd(fr *frames.Frame, _ int64) int {
	rhs := frames.AsFloat64(pop(fr))
	lhs := frames.AsFloat64(pop(fr))
	push(fr, frames.Float64Slot(lhs+rhs))
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
	var slot frames.Slot
	var emptyOpStack []frames.Slot
	var fr frames.Frame
	var pi64 = 3.14
	var pi64p1 = 4.14
	var one32 float32 = 1.0
	var one64 = 1.0
	var pi32 = float32(3.14)

	for range 10 {
		emptyOpStack = append(emptyOpStack, frames.IntSlot(0))
	}
	var locals []frames.Slot

	for ix := range numLoops {

		fr = frames.Frame{OpStack: emptyOpStack, Locals: locals, TOS: -1}

		// doIadd
		push(&fr, frames.IntSlot(math.MaxInt32-1))
		push(&fr, frames.IntSlot(1))
		doIadd(&fr, 0)
		slot = pop(&fr)

		// doLadd
		push(&fr, frames.IntSlot(math.MaxInt64-1))
		push(&fr, frames.IntSlot(1))
		doLadd(&fr, 0)
		slot = pop(&fr)

		// doFadd
		push(&fr, frames.Float32Slot(pi32))
		push(&fr, frames.Float32Slot(one32))
		doFadd(&fr, 0)
		slot = pop(&fr)

		// doDadd
		push(&fr, frames.Float64Slot(pi64))
		push(&fr, frames.Float64Slot(one64))
		doDadd(&fr, 0)
		slot = pop(&fr)

		if ix&0x1 == 0 {
			// doDup
			push(&fr, frames.Float64Slot(pi64))
			doDup(&fr, 0)
			slot = pop(&fr)
			slot = pop(&fr)

			// doDup2
			push(&fr, frames.Float64Slot(pi64))
			push(&fr, frames.Float64Slot(pi64p1))
			doDup2(&fr, 0)
			slot = pop(&fr)
			slot = pop(&fr)
			slot = pop(&fr)
			slot = pop(&fr)

			// G function
			var thing any
			thing = gfunc([]any{thing})
		}

		// Object
		obj := &object.Object{X: 1, Y: 2, Z: 3}
		push(&fr, frames.ObjSlot(obj))
		slot = pop(&fr)
		obj = frames.AsObject(slot)

	} // for ix := range numLoops

	// Success!
	return true
}

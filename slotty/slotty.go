package slotty

import (
	"bakeoff/slotty/frames"
	"bakeoff/slotty/object"
	"bakeoff/slotty/pslot"
	"fmt"
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

func check(label string, ix int, expected any, observed any) bool {
	var equal bool
	switch observed.(type) {
	case int32:
		exp := expected.(int32)
		obs := observed.(int32)
		equal = exp == obs
	case int64:
		exp := expected.(int64)
		obs := observed.(int64)
		equal = exp == obs
	case float32:
		exp := expected.(float32)
		obs := observed.(float32)
		equal = fmt.Sprintf("%.6f", exp) == fmt.Sprintf("%.6f", obs)
	case float64:
		exp := expected.(float64)
		obs := observed.(float64)
		equal = fmt.Sprintf("%.6f", exp) == fmt.Sprintf("%.6f", obs)
	default:
		panic(fmt.Sprintf("Unexpected type for observed: %T", observed))
	}
	if !equal {
		fmt.Printf("*** ERROR, %s [%d] expected: %v, observed: %v\n", label, ix, expected, observed)
		return false
	}
	return true
}

func Slotty(argNloops int) bool {
	var slot pslot.Slot
	var emptyOpStack []pslot.Slot
	var fr frames.Frame
	var pi64 = 3.14
	var pi64p1 = 4.14
	var one32 float32 = 1.0
	var one64 = 1.0

	for range 10 {
		emptyOpStack = append(emptyOpStack, pslot.IntSlot(0))
	}
	var locals []pslot.Slot

	// doIadd
	fr = frames.Frame{OpStack: emptyOpStack, Locals: locals, TOS: -1}
	for ix := range argNloops {
		push(&fr, pslot.IntSlot(math.MaxInt32-1))
		push(&fr, pslot.IntSlot(1))
		doIadd(&fr, 0)
		slot = pop(&fr)
		if !check("doIadd", ix, int32(math.MaxInt32), pslot.AsInt32(slot)) {
			return false
		}
	}

	// doLadd
	fr = frames.Frame{OpStack: emptyOpStack, Locals: locals, TOS: -1}
	for ix := range argNloops {
		push(&fr, pslot.IntSlot(math.MaxInt64-1))
		push(&fr, pslot.IntSlot(1))
		doLadd(&fr, 0)
		slot = pop(&fr)
		if !check("doLadd", ix, int64(math.MaxInt64), pslot.AsInt64(slot)) {
			return false
		}
	}

	// doFadd
	fr = frames.Frame{OpStack: emptyOpStack, Locals: locals, TOS: -1}
	var pi32 = float32(3.14)
	var pi32p1 = float32(4.14)
	for ix := range argNloops {
		push(&fr, pslot.Float32Slot(pi32))
		push(&fr, pslot.Float32Slot(one32))
		doFadd(&fr, 0)
		slot = pop(&fr)
		if !check("doFadd", ix, pi32p1, pslot.AsFloat32(slot)) {
			return false
		}
	}

	// doDadd
	fr = frames.Frame{OpStack: emptyOpStack, Locals: locals, TOS: -1}
	for ix := range argNloops {
		push(&fr, pslot.Float64Slot(pi64))
		push(&fr, pslot.Float64Slot(one64))
		doDadd(&fr, 0)
		slot = pop(&fr)
		if !check("doDadd", ix, pi64p1, pslot.AsFloat64(slot)) {
			return false
		}
	}

	// doDup
	fr = frames.Frame{OpStack: emptyOpStack, Locals: locals, TOS: -1}
	for ix := range argNloops {
		push(&fr, pslot.Float64Slot(pi64))
		doDup(&fr, 0)
		slot = pop(&fr)
		if !check("doDup top", ix, pi64, pslot.AsFloat64(slot)) {
			return false
		}
		slot = pop(&fr)
		if !check("doDup next", ix, pi64, pslot.AsFloat64(slot)) {
			return false
		}
	}

	// doDup2
	fr = frames.Frame{OpStack: emptyOpStack, Locals: locals, TOS: -1}
	for ix := range argNloops {
		push(&fr, pslot.Float64Slot(pi64))
		push(&fr, pslot.Float64Slot(pi64p1))
		doDup2(&fr, 0)
		slot = pop(&fr)
		if !check("doDup2 top", ix, pi64p1, pslot.AsFloat64(slot)) {
			return false
		}
		slot = pop(&fr)
		if !check("doDup2 #2", ix, pi64, pslot.AsFloat64(slot)) {
			return false
		}
		slot = pop(&fr)
		if !check("doDup2 #3", ix, pi64p1, pslot.AsFloat64(slot)) {
			return false
		}
		slot = pop(&fr)
		if !check("doDup2 #4", ix, pi64, pslot.AsFloat64(slot)) {
			return false
		}
	}

	// Object
	fr = frames.Frame{OpStack: emptyOpStack, Locals: locals, TOS: -1}
	obj := &object.Object{X: 1, Y: 2, Z: 3}
	for range argNloops {
		push(&fr, pslot.ObjSlot(obj))
		slot = pop(&fr)
		obj = pslot.AsObject(slot)
	}
	if obj.X != 1 || obj.Y != 2 || obj.Z != 3 {
		return false
	}

	return true
}

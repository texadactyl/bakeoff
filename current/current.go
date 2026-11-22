package current

import (
	"bakeoff/current/frames"
	"bakeoff/current/object"
	"fmt"
	"math"
)

func push(fr *frames.Frame, thing any) {
	fr.TOS++
	fr.OpStack[fr.TOS] = thing
}

func pop(fr *frames.Frame) any {
	thing := fr.OpStack[fr.TOS]
	fr.TOS--
	return thing
}

func doIadd(fr *frames.Frame, _ int64) int {
	i2 := pop(fr).(int64)
	i1 := pop(fr).(int64)
	sum := int32(i1) + int32(i2)
	push(fr, int64(sum))
	return 1
}

func doLadd(fr *frames.Frame, _ int64) int {
	i2 := pop(fr).(int64)
	i1 := pop(fr).(int64)
	push(fr, i1+i2)
	return 1
}

func doFadd(fr *frames.Frame, _ int64) int {
	rhs := float32(pop(fr).(float64))
	lhs := float32(pop(fr).(float64))
	push(fr, float64(lhs+rhs))
	return 1
}

func doDadd(fr *frames.Frame, _ int64) int {
	rhs := pop(fr).(float64)
	lhs := pop(fr).(float64)
	push(fr, lhs+rhs)
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

func Current(NLOOPS int) bool {
	var thing any
	var emptyOpStack []any
	var fr frames.Frame
	for range 10 {
		emptyOpStack = append(emptyOpStack, int64(0))
	}
	var locals []any

	// doIadd
	fr = frames.Frame{OpStack: emptyOpStack, Locals: locals, TOS: -1}
	for ix := 0; ix < NLOOPS; ix++ {
		push(&fr, int64(math.MaxInt32-1))
		push(&fr, int64(1))
		doIadd(&fr, 0)
		thing = pop(&fr)
		if !check("doIadd", ix, int32(math.MaxInt32), int32(thing.(int64))) {
			return false
		}
	}

	// doLadd
	fr = frames.Frame{OpStack: emptyOpStack, Locals: locals, TOS: -1}
	for ix := 0; ix < NLOOPS; ix++ {
		push(&fr, int64(math.MaxInt64-1))
		push(&fr, int64(1))
		doLadd(&fr, 0)
		thing = pop(&fr)
		if !check("doLadd", ix, int64(math.MaxInt64), thing.(int64)) {
			return false
		}
	}

	// doFadd
	fr = frames.Frame{OpStack: emptyOpStack, Locals: locals, TOS: -1}
	var pi32 = float32(3.14)
	var pi32p1 = float32(4.14)
	for ix := 0; ix < NLOOPS; ix++ {
		push(&fr, float64(pi32))
		push(&fr, float64(1.0))
		doFadd(&fr, 0)
		thing = pop(&fr)
		if !check("doFadd", ix, pi32p1, float32(thing.(float64))) {
			return false
		}
	}

	// doDadd
	fr = frames.Frame{OpStack: emptyOpStack, Locals: locals, TOS: -1}
	var pi64 = float64(3.14)
	var pi64p1 = float64(4.14)
	for ix := 0; ix < NLOOPS; ix++ {
		push(&fr, pi64)
		push(&fr, float64(1.0))
		doDadd(&fr, 0)
		thing = pop(&fr)
		if !check("doDadd", ix, pi64p1, thing.(float64)) {
			return false
		}
	}

	// Object
	fr = frames.Frame{OpStack: emptyOpStack, Locals: locals, TOS: -1}
	obj := &object.Object{X: 1, Y: 2, Z: 3}
	for range NLOOPS {
		push(&fr, obj)
		obj = pop(&fr).(*object.Object)
	}
	if obj.X != 1 || obj.Y != 2 || obj.Z != 3 {
		return false
	}

	return true
}

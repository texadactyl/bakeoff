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

func peek(fr *frames.Frame) any {
	return fr.OpStack[fr.TOS]
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

func Current(numLoops int, dupDup2Modulus int) bool {
	var thing any
	var emptyOpStack []any
	var fr frames.Frame
	var pi64 = 3.14
	var pi64p1 = 4.14
	var one = 1.0
	var pi32 = float32(3.14)
	var pi32p1 = float32(4.14)

	for range 10 {
		emptyOpStack = append(emptyOpStack, int64(0))
	}
	var locals []any

	for ix := range numLoops {
	
		// doIadd
		fr = frames.Frame{OpStack: emptyOpStack, Locals: locals, TOS: -1}
		push(&fr, int64(math.MaxInt32-1))
		push(&fr, int64(1))
		doIadd(&fr, 0)
		thing = pop(&fr)
		if !check("doIadd", ix, int32(math.MaxInt32), int32(thing.(int64))) {
			return false
		}
		
		// doLadd
		fr = frames.Frame{OpStack: emptyOpStack, Locals: locals, TOS: -1}
		push(&fr, int64(math.MaxInt64-1))
		push(&fr, int64(1))
		doLadd(&fr, 0)
		thing = pop(&fr)
		if !check("doLadd", ix, int64(math.MaxInt64), thing.(int64)) {
			return false
		}

		// doFadd
		fr = frames.Frame{OpStack: emptyOpStack, Locals: locals, TOS: -1}
		push(&fr, float64(pi32))
		push(&fr, one)
		doFadd(&fr, 0)
		thing = pop(&fr)
		if !check("doFadd", ix, pi32p1, float32(thing.(float64))) {
			return false
		}

		// doDadd
		fr = frames.Frame{OpStack: emptyOpStack, Locals: locals, TOS: -1}
		push(&fr, pi64)
		push(&fr, one)
		doDadd(&fr, 0)
		thing = pop(&fr)
		if !check("doDadd", ix, pi64p1, thing.(float64)) {
			return false
		}

		if ix%dupDup2Modulus == 0 {
		
			// doDup
			fr = frames.Frame{OpStack: emptyOpStack, Locals: locals, TOS: -1}
			push(&fr, pi64)
			doDup(&fr, 0)
			thing = pop(&fr)
			if !check("doDup top", ix, pi64, thing.(float64)) {
				return false
			}
			thing = pop(&fr)
			if !check("doDup next", ix, pi64, thing.(float64)) {
				return false
			}

			// doDup2
			fr = frames.Frame{OpStack: emptyOpStack, Locals: locals, TOS: -1}
			push(&fr, pi64)
			push(&fr, pi64p1)
			doDup2(&fr, 0)
			thing = pop(&fr)
			if !check("doDup2 top", ix, pi64p1, thing.(float64)) {
				return false
			}
			thing = pop(&fr)
			if !check("doDup2 #2", ix, pi64, thing.(float64)) {
				return false
			}
			thing = pop(&fr)
			if !check("doDup2 #3", ix, pi64p1, thing.(float64)) {
				return false
			}
			thing = pop(&fr)
			if !check("doDup2 #4", ix, pi64, thing.(float64)) {
				return false
			}

		}

		// Object
		fr = frames.Frame{OpStack: emptyOpStack, Locals: locals, TOS: -1}
		obj := &object.Object{X: 1, Y: 2, Z: 3}
		push(&fr, obj)
		obj = pop(&fr).(*object.Object)
		if obj.X != 1 || obj.Y != 2 || obj.Z != 3 {
			return false
		}

	} // for ix := range numLoops

	// Success!
	return true

}

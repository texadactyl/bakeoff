package current

import (
	"bakeoff/current/frames"
	"bakeoff/current/object"
	"math"
)

func gfunc(params[]any) any {
	return int64(0)
}

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

func Current(numLoops int) bool {
	var thing any
	var emptyOpStack []any
	var fr frames.Frame
	var pi64 = 3.14
	var pi64p1 = 4.14
	var one = 1.0
	var pi32 = float32(3.14)

	for range 10 {
		emptyOpStack = append(emptyOpStack, int64(0))
	}
	var locals []any

	for ix := range numLoops {

		fr = frames.Frame{OpStack: emptyOpStack, Locals: locals, TOS: -1}

		// doIadd
		push(&fr, int64(math.MaxInt32-1))
		push(&fr, int64(1))
		doIadd(&fr, 0)
		thing = pop(&fr)

		// doLadd
		push(&fr, int64(math.MaxInt64-1))
		push(&fr, int64(1))
		doLadd(&fr, 0)
		thing = pop(&fr)

		// doFadd
		push(&fr, float64(pi32))
		push(&fr, one)
		doFadd(&fr, 0)
		thing = pop(&fr)

		// doDadd
		push(&fr, pi64)
		push(&fr, one)
		doDadd(&fr, 0)
		thing = pop(&fr)

		if ix&0x1 == 0 {

			// doDup
			push(&fr, pi64)
			doDup(&fr, 0)
			thing = pop(&fr)
			thing = pop(&fr)

			// doDup2
			push(&fr, pi64)
			push(&fr, pi64p1)
			doDup2(&fr, 0)
			thing = pop(&fr)
			thing = pop(&fr)
			thing = pop(&fr)
			thing = pop(&fr)

			// G function
			thing = gfunc([]any{thing})

		}

		// Object
		thing = &object.Object{X: 1, Y: 2, Z: 3}
		push(&fr, thing)
		thing = pop(&fr).(*object.Object)

	} // for ix := range numLoops

	// Success!
	return true

}

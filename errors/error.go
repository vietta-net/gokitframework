// Copyright (c) 2016 The Upspin Authors. All rights reserved.
// A lot of this was stolen from upspin
// (https://github.com/upspin/upspin/blob/master/errors/errors.go) and then
// modified.

package errors

// populateStack uses the runtime to populate the Error's stack struct with
import (
	"github.com/vietta-net/agokit/helper"
	"bytes"
	"fmt"
	"google.golang.org/grpc/codes"
	"runtime"
	"strings"
)


// assert Error implements the error interface.
var (
	_ error = &Error{}
	_ helper.Copy = &Error{}
)

func (e *Error) To(data interface{}) (err error)  {
	err = helper.Convert(e, &data)
	return err
}

func (e *Error) From(data interface{}) (err error)  {
	err = helper.Convert(data, &e)
	return err
}

func (e *Error) SetCode(code codes.Code)  {
	e.Code = uint32(code)
}



// Error implements the error interface.
func (e *Error) Error() string {
	b := new(bytes.Buffer)
	e.printStack(b)
	pad(b, ": ")
	b.WriteString(e.Message)
	if b.Len() == 0 {
		return "no error"
	}
	return b.String()
}

// WrapErr returns a corev1.Error for the given error and msg.
func WrapErr(err error, msg string) error {
	if err == nil {
		return nil
	}
	e := &Error{Message: fmt.Sprintf("%s: %s", msg, err.Error())}
	e.populateStack()
	return e
}

// E is a useful func for instantiating corev1.Errors.
func E(args ...interface{}) error {
	if len(args) == 0 {
		panic("call to E with no arguments")
	}
	e := &Error{}
	b := new(bytes.Buffer)
	for _, arg := range args {
		switch arg := arg.(type) {
		case string:
			pad(b, ": ")
			b.WriteString(arg)
		case codes.Code:
			e.Code = uint32(arg)
		case map[string]string:
			e.NestedErrors = arg
		case int32:
			e.Code = uint32(arg)
		case int:
			e.Code = uint32(int32(arg))
		case error:
			pad(b, ": ")
			b.WriteString(arg.Error())
		}
	}
	e.Message = b.String()
	e.populateStack()
	return e
}

// populateStack uses the runtime to populate the Error's stack struct with
// information about the current stack.
func (e *Error) populateStack() {
	var s = &Stack{}
	s.Callers = callers()
	e.Stack, _ = s.Marshal()
}

// printStack formats and prints the stack for this Error to the given buffer.
// It should be called from the Error's Error method.
func (e *Error) printStack(b *bytes.Buffer) {
	if e.Stack == nil {
		return
	}
	var stack = &Stack{}
	stack.Unmarshal(e.Stack)

	printCallers := callers()

	// Iterate backward through e.Stack.Callers (the last in the stack is the
	// earliest call, such as main) skipping over the PCs that are shared
	// by the error stack and by this function call stack, printing the
	// names of the functions and their file names and line numbers.
	var prev string // the name of the last-seen function
	var diff bool   // do the print and error call stacks differ now?
	for i := 0; i < len(stack.Callers); i++ {
		thisFrame := frame(stack.Callers, i)
		name := thisFrame.Func.Name()

		if !diff && i < len(printCallers) {
			if name == frame(printCallers, i).Func.Name() {
				// both stacks share this PC, skip it.
				continue
			}
			// No match, don't consider printCallers again.
			diff = true
		}

		// Don't print the same function twice.
		// (Can happen when multiple error stacks have been coalesced.)
		if name == prev {
			continue
		}

		// Find the uncommon prefix between this and the previous
		// function name, separating by dots and slashes.
		trim := 0
		for {
			j := strings.IndexAny(name[trim:], "./")
			if j < 0 {
				break
			}
			if !strings.HasPrefix(prev, name[:j+trim]) {
				break
			}
			trim += j + 1 // skip over the separator
		}

		// Do the printing.
		pad(b, separator)
		fmt.Fprintf(b, "%v:%d: ", thisFrame.File, thisFrame.Line)
		if trim > 0 {
			b.WriteString("...")
		}
		b.WriteString(name[trim:])

		prev = name
	}
}

// frame returns the nth frame, with the frame at top of stack being 0.
func frame(callers []uintptr, n int) *runtime.Frame {
	frames := runtime.CallersFrames(callers)
	var f runtime.Frame
	for i := len(callers) - 1; i >= n; i-- {
		var ok bool
		f, ok = frames.Next()
		if !ok {
			break // Should never happen, and this is just debugging.
		}
	}
	return &f
}

// callers is a wrapper for runtime.callers that allocates a slice.
func callers() []uintptr {
	var stk [64]uintptr
	const skip = 4 // Skip 4 stack frames; ok for both E and Error funcs.
	n := runtime.Callers(skip, stk[:])
	return stk[:n]
}

var separator = ":\n\t"

// pad appends str to the buffer if the buffer already has some data.
func pad(b *bytes.Buffer, str string) {
	if b.Len() == 0 {
		return
	}
	b.WriteString(str)
}


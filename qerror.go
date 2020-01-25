package qerror

import (
	"fmt"
	"runtime"
)

type ErrorID uint32

var messages map[ErrorID]string

func Init(m map[ErrorID]string) {
	messages = m
}

type Error struct {
	ErrorID    int64
	Args       []interface{}
	StackFile  string
	StackLine  int
	StackFunc  string
	InnerError error
}

func (e *Error) Message() string {
	if e.ErrorID < 0 {
		return ""
	}

	if m, ok := messages[ErrorID(e.ErrorID)]; ok {
		return fmt.Sprintf(m, e.Args...)
	} else {
		return fmt.Sprintf("%+v", e.Args)
	}
}

func (e *Error) Error() string {
	msg := fmt.Sprintf("%s\n\t%s:%d %s",
		e.Message(),
		e.StackFile,
		e.StackLine,
		e.StackFunc,
	)

	if e.InnerError == nil {
		return msg
	}

	return fmt.Sprintf("%s\ncaused by: %s", msg, e.InnerError.Error())
}

func IsQError(err error) bool {
	_, ok := err.(*Error)
	return ok
}

func New(errorID ErrorID, args ...interface{}) error {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "UNKNOWN"
		line = -1
	}
	fn := runtime.FuncForPC(pc)
	fnName := "UNKNOWN FUNC"
	if fn != nil {
		fnName = fn.Name()
	}

	return &Error{
		ErrorID:    int64(errorID),
		Args:       args,
		StackFile:  file,
		StackLine:  line,
		StackFunc:  fnName,
		InnerError: nil,
	}
}

func Wrap(e error) error {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "UNKNOWN"
		line = -1
	}
	fn := runtime.FuncForPC(pc)
	fnName := "UNKNOWN FUNC"
	if fn != nil {
		fnName = fn.Name()
	}

	return &Error{
		ErrorID:    -1,
		Args:       nil,
		StackFile:  file,
		StackLine:  line,
		StackFunc:  fnName,
		InnerError: e,
	}
}

func WrapWith(e error, errorID ErrorID, args ...interface{}) error {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "UNKNOWN"
		line = -1
	}
	fn := runtime.FuncForPC(pc)
	fnName := "UNKNOWN FUNC"
	if fn != nil {
		fnName = fn.Name()
	}

	return &Error{
		ErrorID:    int64(errorID),
		Args:       args,
		StackFile:  file,
		StackLine:  line,
		StackFunc:  fnName,
		InnerError: e,
	}
}

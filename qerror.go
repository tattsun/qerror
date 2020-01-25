package qerror

import (
	"fmt"
	"runtime"
)

var messages map[uint32]string

func Init(m map[uint32]string) {
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

	if m, ok := messages[uint32(e.ErrorID)]; ok {
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

func New(errorID int64, args ...interface{}) error {
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
		ErrorID:    errorID,
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

func WrapWith(e error, errorID int64, args ...interface{}) error {
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
		ErrorID:    errorID,
		Args:       args,
		StackFile:  file,
		StackLine:  line,
		StackFunc:  fnName,
		InnerError: e,
	}
}

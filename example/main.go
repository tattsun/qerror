package main

import (
	"fmt"

	"github.com/tattsun/qerror"
)

func testA() error {
	return qerror.New(InvalidArgument, "hoge")
}

func testB() error {
	err := testA()
	return qerror.WrapWith(err, DBError)
}

func testC() error {
	err := testB()
	return qerror.WrapWith(err, ValidationError, "fuga", "piyo")
}

const (
	InvalidArgument qerror.ErrorID = iota
	DBError
	ValidationError
)

func main() {
	qerror.Init(map[qerror.ErrorID]string{
		InvalidArgument: "invalid argument: %s",
		DBError:         "db error",
		ValidationError: "validation error: %s %s",
	})

	fmt.Println(testC())
}

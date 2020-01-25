package main

import (
	"fmt"

	"github.com/tattsun/qerror"
)

func testA() error {
	return qerror.New(1, "hoge")
}

func testB() error {
	err := testA()
	return qerror.WrapWith(err, 2)
}

func testC() error {
	err := testB()
	return qerror.WrapWith(err, 3, "fuga", "piyo")
}

func main() {
	qerror.Init(map[uint32]string{
		1: "invalid argument: %s",
		2: "db error",
		3: "validation error: %s %s",
	})

	fmt.Println(testC())
}

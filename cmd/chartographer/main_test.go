package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFunctions(t *testing.T) {
	Convey("Testing `getGreeting` function", t, func() {
		So(getGreeting(), ShouldEqual, "Hello, Kontur!")
	})
}

package vo

import (
	"account-management-service/internal/core/failure"
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFullName(t *testing.T) {
	Convey("Test full name", t, func() {
		Convey("Empty full name", func() {
			const fullName = ""

			f, err := NewFullName(fullName)

			So(errors.Is(err, failure.EmptyFullNameError{}), ShouldBeTrue)
			So(err.Error(), ShouldEqual, failure.EmptyFullNameError{}.Error())
			So(f, ShouldBeNil)
		})

		Convey("Minimal valid full name", func() {
			const fullName = "a"

			f, err := NewFullName(fullName)

			So(err, ShouldBeNil)
			So(f.Value(), ShouldEqual, fullName)
		})

		Convey("Normal full name", func() {
			const fullName = "Alexey Rachkov"

			f, err := NewFullName(fullName)

			So(err, ShouldBeNil)
			So(f.Value(), ShouldEqual, fullName)
		})
	})
}

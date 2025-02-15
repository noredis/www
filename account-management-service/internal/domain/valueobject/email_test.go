package vo

import (
	"account-management-service/internal/core/failure"
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEmail(t *testing.T) {
	Convey("Test email", t, func() {
		Convey("Empty email", func() {
			const email = ""

			e, err := NewEmail(email)

			So(errors.Is(err, failure.EmptyEmailError{}), ShouldBeTrue)
			So(err.Error(), ShouldEqual, failure.EmptyEmailError{}.Error())
			So(e, ShouldBeNil)
		})

		Convey("Invalid email", func() {
			const email = "alexey"

			e, err := NewEmail(email)

			So(errors.Is(err, failure.InvalidEmailError{}), ShouldBeTrue)
			So(err.Error(), ShouldEqual, failure.InvalidEmailError{}.Error())
			So(e, ShouldBeNil)
		})

		Convey("Normal email", func() {
			const email = "alexey@gmail.com"

			e, err := NewEmail(email)

			So(err, ShouldBeNil)
			So(e.Value(), ShouldEqual, email)
		})
	})
}

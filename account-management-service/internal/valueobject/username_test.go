package vo

import (
	"account-management-service/internal/failure"
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUsername(t *testing.T) {
	Convey("Test username", t, func() {
		Convey("Empty username", func() {
			const username = ""

			u, err := NewUsername(username)

			So(errors.Is(err, failure.EmptyUsernameError{}), ShouldBeTrue)
			So(err.Error(), ShouldEqual, failure.EmptyUsernameError{}.Error())
			So(u, ShouldBeNil)
		})

		Convey("Too short username", func() {
			const username = "mira"

			u, err := NewUsername(username)

			So(errors.Is(err, failure.UsernameTooShortError{}), ShouldBeTrue)
			So(err.Error(), ShouldEqual, failure.UsernameTooShortError{}.Error())
			So(u, ShouldBeNil)
		})

		Convey("Minimal length username", func() {
			const username = "pirat"

			u, err := NewUsername(username)

			So(err, ShouldBeNil)
			So(u.Value(), ShouldEqual, username)
		})

		Convey("Too long username", func() {
			const username = "piratpiratpiratpiratp"

			u, err := NewUsername(username)

			So(errors.Is(err, failure.UsernameTooLongError{}), ShouldBeTrue)
			So(err.Error(), ShouldEqual, failure.UsernameTooLongError{}.Error())
			So(u, ShouldBeNil)
		})

		Convey("Maximum length username", func() {
			const username = "piratpiratpiratpirat"

			u, err := NewUsername(username)

			So(err, ShouldBeNil)
			So(u.Value(), ShouldEqual, username)
		})

		Convey("Invalid username", func() {
			const username = "bruh lord"

			u, err := NewUsername(username)

			So(errors.Is(err, failure.InvalidUsernameError{}), ShouldBeTrue)
			So(err.Error(), ShouldEqual, failure.InvalidUsernameError{}.Error())
			So(u, ShouldBeNil)
		})

		Convey("Normal username", func() {
			const username = "killoverm"

			u, err := NewUsername(username)

			So(err, ShouldBeNil)
			So(u.Value(), ShouldEqual, username)
		})

		Convey("Normal username with number", func() {
			const username = "vispers0"

			u, err := NewUsername(username)

			So(err, ShouldBeNil)
			So(u.Value(), ShouldEqual, username)
		})

		Convey("Normal username with underscore", func() {
			const username = "chulochki_prime"

			u, err := NewUsername(username)

			So(err, ShouldBeNil)
			So(u.Value(), ShouldEqual, username)
		})

		Convey("Normal username with dot", func() {
			const username = "hard.lime"

			u, err := NewUsername(username)

			So(err, ShouldBeNil)
			So(u.Value(), ShouldEqual, username)
		})
	})
}

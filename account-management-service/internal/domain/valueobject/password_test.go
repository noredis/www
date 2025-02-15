package vo

import (
	"account-management-service/internal/core/failure"
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPassword(t *testing.T) {
	Convey("Test password", t, func() {
		Convey("Empty password", func() {
			const (
				password             = ""
				passwordConfirmation = password
			)

			p, err := NewPassword(password, passwordConfirmation)

			So(errors.Is(err, failure.EmptyPasswordError{}), ShouldBeTrue)
			So(err.Error(), ShouldEqual, failure.EmptyPasswordError{}.Error())
			So(p, ShouldBeNil)
		})

		Convey("Too short password", func() {
			const (
				password             = "p5sW$"
				passwordConfirmation = password
			)

			p, err := NewPassword(password, passwordConfirmation)

			So(errors.Is(err, failure.PasswordTooShortError{}), ShouldBeTrue)
			So(err.Error(), ShouldEqual, failure.PasswordTooShortError{}.Error())
			So(p, ShouldBeNil)
		})

		Convey("Minimal password length", func() {
			const (
				password             = "p5sW$r"
				passwordConfirmation = password
			)

			p, err := NewPassword(password, passwordConfirmation)

			So(err, ShouldBeNil)
			So(p.Value(), ShouldNotEqual, password)
			So(p.Compare(password), ShouldBeTrue)
		})

		Convey("Too long password", func() {
			const (
				password             = "pa5sW$rdpa5sW$rdpa5sW$rdpa5sW$rdpa5sW$rdp"
				passwordConfirmation = password
			)

			p, err := NewPassword(password, passwordConfirmation)

			So(errors.Is(err, failure.PasswordTooLongError{}), ShouldBeTrue)
			So(err.Error(), ShouldEqual, failure.PasswordTooLongError{}.Error())
			So(p, ShouldBeNil)
		})

		Convey("Maximum password length", func() {
			const (
				password             = "pa5sW$rdpa5sW$rdpa5sW$rdpa5sW$rdpa5sW$rd"
				passwordConfirmation = password
			)

			p, err := NewPassword(password, passwordConfirmation)

			So(err, ShouldBeNil)
			So(p.Value(), ShouldNotEqual, password)
			So(p.Compare(password), ShouldBeTrue)
		})

		Convey("Password mismatch", func() {
			const (
				password             = "pa5sW$rd"
				passwordConfirmation = "passW0$d"
			)

			p, err := NewPassword(password, passwordConfirmation)

			So(errors.Is(err, failure.PasswordMismatchError{}), ShouldBeTrue)
			So(err.Error(), ShouldEqual, failure.PasswordMismatchError{}.Error())
			So(p, ShouldBeNil)
		})

		Convey("Password without number", func() {
			const (
				password             = "passW$rd"
				passwordConfirmation = password
			)

			p, err := NewPassword(password, passwordConfirmation)

			So(errors.Is(err, failure.InvalidPasswordError{}), ShouldBeTrue)
			So(err.Error(), ShouldEqual, failure.InvalidPasswordError{}.Error())
			So(p, ShouldBeNil)
		})

		Convey("Password without uppercase letter", func() {
			const (
				password             = "pa5sw$rd"
				passwordConfirmation = password
			)

			p, err := NewPassword(password, passwordConfirmation)

			So(errors.Is(err, failure.InvalidPasswordError{}), ShouldBeTrue)
			So(err.Error(), ShouldEqual, failure.InvalidPasswordError{}.Error())
			So(p, ShouldBeNil)
		})

		Convey("Password without lowercase letter", func() {
			const (
				password             = "PA5SW$RD"
				passwordConfirmation = password
			)

			p, err := NewPassword(password, passwordConfirmation)

			So(errors.Is(err, failure.InvalidPasswordError{}), ShouldBeTrue)
			So(err.Error(), ShouldEqual, failure.InvalidPasswordError{}.Error())
			So(p, ShouldBeNil)
		})

		Convey("Password without special characters", func() {
			const (
				password             = "pa5sWord"
				passwordConfirmation = password
			)

			p, err := NewPassword(password, passwordConfirmation)

			So(errors.Is(err, failure.InvalidPasswordError{}), ShouldBeTrue)
			So(err.Error(), ShouldEqual, failure.InvalidPasswordError{}.Error())
			So(p, ShouldBeNil)
		})

		Convey("Normal password", func() {
			const (
				password             = "Giv3MeF!ve"
				passwordConfirmation = password
			)

			p, err := NewPassword(password, passwordConfirmation)

			So(err, ShouldBeNil)
			So(p.Value(), ShouldNotEqual, password)
			So(p.Compare(password), ShouldBeTrue)

			p, err = RestorePassword(p.Value())

			So(err, ShouldBeNil)
			So(p.Value(), ShouldNotEqual, password)
		})

		Convey("Restore not hashed password", func() {
			const password = "Giv3MeF!ve"
			p, err := RestorePassword(password)

			So(errors.Is(err, failure.UnableToRestorePasswordError{}), ShouldBeTrue)
			So(err.Error(), ShouldEqual, failure.UnableToRestorePasswordError{}.Error())
			So(p, ShouldBeNil)
		})
	})
}

package entity

import (
	vo "account-management-service/internal/valueobject"
	"github.com/google/uuid"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAccount(t *testing.T) {
	Convey("Test account", t, func() {
		Convey("Create new account", func() {
			id, err := uuid.NewUUID()
			So(err, ShouldBeNil)

			const f = "Alexey Rachkov"
			fullName, err := vo.NewFullName(f)
			So(err, ShouldBeNil)

			const e = "alexey.rachkov@gmail.com"
			email, err := vo.NewEmail(e)
			So(err, ShouldBeNil)

			const u = "a.rachkov"
			username, err := vo.NewUsername(u)
			So(err, ShouldBeNil)

			const p = "pas5W$rd"
			password, err := vo.NewPassword(p, p)
			So(err, ShouldBeNil)

			now := time.Now()

			account := Register(id, *fullName, *email, *username, *password, now)

			So(account.ID(), ShouldEqual, id)
			So(account.FullName().Value(), ShouldEqual, f)
			So(account.Email().Value(), ShouldEqual, e)
			So(account.Username().Value(), ShouldEqual, u)
			So(account.Password().Compare(p), ShouldBeTrue)
			So(account.CreatedAt(), ShouldEqual, now)
			So(account.PasswordUpdatedAt(), ShouldEqual, now)
			So(account.EmailConfirmedAt(), ShouldBeNil)

			account = RestoreAccount(
				account.ID(),
				account.FullName(),
				account.Email(),
				account.Username(),
				account.Password(),
				account.CreatedAt(),
				account.PasswordUpdatedAt(),
				account.EmailConfirmedAt(),
			)

			So(account.ID(), ShouldEqual, id)
			So(account.FullName().Value(), ShouldEqual, f)
			So(account.Email().Value(), ShouldEqual, e)
			So(account.Username().Value(), ShouldEqual, u)
			So(account.Password().Compare(p), ShouldBeTrue)
			So(account.CreatedAt(), ShouldEqual, now)
			So(account.PasswordUpdatedAt(), ShouldEqual, now)
			So(account.EmailConfirmedAt(), ShouldBeNil)
		})
	})
}

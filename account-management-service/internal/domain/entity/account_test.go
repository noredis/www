package entity

import (
	domainevent2 "account-management-service/internal/domain/domainevent"
	vo2 "account-management-service/internal/domain/valueobject"
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
			fullName, err := vo2.NewFullName(f)
			So(err, ShouldBeNil)

			const e = "alexey.rachkov@gmail.com"
			email, err := vo2.NewEmail(e)
			So(err, ShouldBeNil)

			const u = "a.rachkov"
			username, err := vo2.NewUsername(u)
			So(err, ShouldBeNil)

			const p = "pas5W$rd"
			password, err := vo2.NewPassword(p, p)
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

			accountEvents := make([]domainevent2.DomainEvent, 0)
			accountEvents = append(accountEvents, domainevent2.NewAccountCreatedEvent(id, *fullName, *email, *username))

			So(account.DomainEvents(), ShouldEqual, accountEvents)

			account.ClearDomainEvents()

			So(account.DomainEvents(), ShouldBeEmpty)

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

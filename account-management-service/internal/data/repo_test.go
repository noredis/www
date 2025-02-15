package data

import (
	"account-management-service/internal/domain/entity"
	vo2 "account-management-service/internal/domain/valueobject"
	"account-management-service/pkg/testingpg"
	"context"
	"github.com/google/uuid"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAccountRepository(t *testing.T) {
	Convey("Test account repository", t, func() {
		Convey("Create new account", func() {
			postgres := testingpg.NewWithIsolatedDatabase(t)
			repo := NewAccountRepository(postgres.DB())

			id, _ := uuid.NewUUID()
			fullName, _ := vo2.NewFullName("Alexey Rachkov")
			email, _ := vo2.NewEmail("alexey.rachkov@gmail.com")
			username, _ := vo2.NewUsername("a.rachkov")
			password, _ := vo2.NewPassword("pas5W$rd", "pas5W$rd")
			now := time.Now()

			account := entity.Register(id, *fullName, *email, *username, *password, now)

			err := repo.Create(context.Background(), *account)

			So(err, ShouldBeNil)

			account, err = repo.GetByID(context.Background(), id)

			So(err, ShouldBeNil)
			So(account.ID(), ShouldEqual, id)

			account, err = repo.GetByEmail(context.Background(), email.Value())

			So(err, ShouldBeNil)
			So(account.ID(), ShouldEqual, id)

			account, err = repo.GetByUsername(context.Background(), username.Value())

			So(err, ShouldBeNil)
			So(account.ID(), ShouldEqual, id)
		})

		Convey("Get non-existing account", func() {
			postgres := testingpg.NewWithIsolatedDatabase(t)
			repo := NewAccountRepository(postgres.DB())

			id, _ := uuid.NewUUID()

			account, err := repo.GetByID(context.Background(), id)

			So(err, ShouldBeNil)
			So(account, ShouldBeNil)

			email := "alexey.rachkov@gmail.com"

			account, err = repo.GetByEmail(context.Background(), email)

			So(err, ShouldBeNil)
			So(account, ShouldBeNil)

			username := "alexey.rachkov"

			account, err = repo.GetByUsername(context.Background(), username)

			So(err, ShouldBeNil)
			So(account, ShouldBeNil)
		})
	})
}

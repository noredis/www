package service

import (
	"account-management-service/internal/data"
	"account-management-service/internal/entity"
	vo "account-management-service/internal/valueobject"
	"account-management-service/pkg/testingpg"
	"context"
	"github.com/google/uuid"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAccountContext(t *testing.T) {
	Convey("Test account context", t, func() {
		Convey("Save changes", func() {
			postgres := testingpg.NewWithIsolatedDatabase(t)
			repo := data.NewAccountRepository(postgres.DB())
			publisher := data.NewFakePublisher()

			accountContext := NewAccountContextImp(repo, publisher)

			id, _ := uuid.NewUUID()
			fullName, _ := vo.NewFullName("Alexey Rachkov")
			email, _ := vo.NewEmail("alexey.rachkov@gmail.com")
			username, _ := vo.NewUsername("a.rachkov")
			password, _ := vo.NewPassword("pas5W$rd", "pas5W$rd")
			now := time.Now()

			account := entity.Register(id, *fullName, *email, *username, *password, now)

			err := accountContext.SaveChanges(context.Background(), *account)

			So(err, ShouldBeNil)
		})
	})
}

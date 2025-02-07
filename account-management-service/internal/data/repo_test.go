package data

import (
	"account-management-service/internal/entity"
	vo "account-management-service/internal/valueobject"
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
			fullName, _ := vo.NewFullName("Alexey Rachkov")
			email, _ := vo.NewEmail("alexey.rachkov@gmail.com")
			username, _ := vo.NewUsername("a.rachkov")
			password, _ := vo.NewPassword("pas5W$rd", "pas5W$rd")
			now := time.Now()

			account := entity.Register(id, *fullName, *email, *username, *password, now)

			err := repo.Create(context.Background(), *account)

			So(err, ShouldBeNil)
		})
	})
}

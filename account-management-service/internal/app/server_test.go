package app

import (
	pb "account-management-service/gen/go/v1/proto"
	"account-management-service/internal/data"
	"account-management-service/internal/service"
	"account-management-service/pkg/testingpg"
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAccountServer(t *testing.T) {
	Convey("Test account server", t, func() {
		Convey("Create account", func() {
			postgres := testingpg.NewWithIsolatedDatabase(t)
			repo := data.NewAccountRepository(postgres.DB())
			publisher := data.NewFakePublisher()
			accountContext := service.NewAccountContextImp(repo, publisher)
			serv := service.NewAccountService(repo, accountContext)

			server := NewAccountServer(*serv)

			account, err := server.CreateAccount(context.Background(), &pb.CreateAccountRequest{
				FullName:             "John Doe",
				Email:                "john@doe.com",
				Username:             "john1",
				Password:             "Qw!123456",
				PasswordConfirmation: "Qw!123456",
			})

			So(err, ShouldBeNil)
			So(account, ShouldNotBeNil)
		})

		Convey("Create empty", func() {
			postgres := testingpg.NewWithIsolatedDatabase(t)
			repo := data.NewAccountRepository(postgres.DB())
			publisher := data.NewFakePublisher()
			accountContext := service.NewAccountContextImp(repo, publisher)
			service := service.NewAccountService(repo, accountContext)

			server := NewAccountServer(*service)

			account, err := server.CreateAccount(context.Background(), &pb.CreateAccountRequest{
				FullName:             "",
				Email:                "",
				Username:             "",
				Password:             "",
				PasswordConfirmation: "",
			})

			So(err, ShouldNotBeNil)
			So(account, ShouldBeNil)
		})
	})
}

package service

import (
	"account-management-service/internal/data"
	"account-management-service/internal/failure"
	"account-management-service/pkg/testingpg"
	"context"
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAccountService(t *testing.T) {
	Convey("Test account service", t, func() {
		Convey("Register account", func() {
			postgres := testingpg.NewWithIsolatedDatabase(t)
			repo := data.NewAccountRepository(postgres.DB())
			publisher := data.NewFakePublisher()
			accountContext := NewAccountContextImp(repo, publisher)

			service := NewAccountService(repo, accountContext)

			dto := CreateAccountDTO{
				FullName:             "John Doe",
				Email:                "john@doe.com",
				Username:             "johndoe",
				Password:             "Qw!123456",
				PasswordConfirmation: "Qw!123456",
			}

			account, err := service.CreateAccount(context.Background(), dto)

			So(err, ShouldBeNil)
			So(account, ShouldNotBeNil)
			So(publisher.Messages(), ShouldNotBeEmpty)

			Convey("Register account with same email", func() {
				dto = CreateAccountDTO{
					FullName:             "John Doe1",
					Email:                "john@doe.com",
					Username:             "johndoe1",
					Password:             "Qw!1234561",
					PasswordConfirmation: "Qw!1234561",
				}

				account, err = service.CreateAccount(context.Background(), dto)

				So(errors.Is(err, failure.EmailIsBusyError{}), ShouldBeTrue)
				So(err.Error(), ShouldEqual, failure.EmailIsBusyError{}.Error())
				So(account, ShouldBeNil)
			})

			Convey("Register account with same username", func() {
				dto = CreateAccountDTO{
					FullName:             "John Doe1",
					Email:                "john1@doe.com",
					Username:             "johndoe",
					Password:             "Qw!1234561",
					PasswordConfirmation: "Qw!1234561",
				}

				account, err = service.CreateAccount(context.Background(), dto)

				So(errors.Is(err, failure.UsernameIsBusyError{}), ShouldBeTrue)
				So(err.Error(), ShouldEqual, failure.UsernameIsBusyError{}.Error())
				So(account, ShouldBeNil)
			})

			Convey("Register account with empty full name", func() {
				postgres := testingpg.NewWithIsolatedDatabase(t)
				repo := data.NewAccountRepository(postgres.DB())
				publisher := data.NewFakePublisher()
				accountContext := NewAccountContextImp(repo, publisher)

				service := NewAccountService(repo, accountContext)

				dto := CreateAccountDTO{
					FullName:             "",
					Email:                "john1@doe.com",
					Username:             "johndoe1",
					Password:             "Qw!123456",
					PasswordConfirmation: "Qw!123456",
				}

				account, err := service.CreateAccount(context.Background(), dto)

				So(err, ShouldNotBeNil)
				So(account, ShouldBeNil)
				So(publisher.Messages(), ShouldBeEmpty)
			})

			Convey("Register account with empty email", func() {
				postgres := testingpg.NewWithIsolatedDatabase(t)
				repo := data.NewAccountRepository(postgres.DB())
				publisher := data.NewFakePublisher()
				accountContext := NewAccountContextImp(repo, publisher)

				service := NewAccountService(repo, accountContext)

				dto := CreateAccountDTO{
					FullName:             "John Doe",
					Email:                "",
					Username:             "johndoe1",
					Password:             "Qw!123456",
					PasswordConfirmation: "Qw!123456",
				}

				account, err := service.CreateAccount(context.Background(), dto)

				So(err, ShouldNotBeNil)
				So(account, ShouldBeNil)
				So(publisher.Messages(), ShouldBeEmpty)
			})

			Convey("Register account with empty username", func() {
				postgres := testingpg.NewWithIsolatedDatabase(t)
				repo := data.NewAccountRepository(postgres.DB())
				publisher := data.NewFakePublisher()
				accountContext := NewAccountContextImp(repo, publisher)

				service := NewAccountService(repo, accountContext)

				dto := CreateAccountDTO{
					FullName:             "John Doe",
					Email:                "john1@doe.com",
					Username:             "",
					Password:             "Qw!123456",
					PasswordConfirmation: "Qw!123456",
				}

				account, err := service.CreateAccount(context.Background(), dto)

				So(err, ShouldNotBeNil)
				So(account, ShouldBeNil)
				So(publisher.Messages(), ShouldBeEmpty)
			})

			Convey("Register account with empty password", func() {
				postgres := testingpg.NewWithIsolatedDatabase(t)
				repo := data.NewAccountRepository(postgres.DB())
				publisher := data.NewFakePublisher()
				accountContext := NewAccountContextImp(repo, publisher)

				service := NewAccountService(repo, accountContext)

				dto := CreateAccountDTO{
					FullName:             "John Doe",
					Email:                "john1@doe.com",
					Username:             "johndoe1",
					Password:             "",
					PasswordConfirmation: "Qw!123456",
				}

				account, err := service.CreateAccount(context.Background(), dto)

				So(err, ShouldNotBeNil)
				So(account, ShouldBeNil)
				So(publisher.Messages(), ShouldBeEmpty)
			})
		})
	})
}

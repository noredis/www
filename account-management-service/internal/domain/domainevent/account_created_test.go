package domainevent

import (
	vo2 "account-management-service/internal/domain/valueobject"
	"github.com/google/uuid"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAccountCreatedEvent(t *testing.T) {
	Convey("Test account created event", t, func() {
		id, _ := uuid.NewUUID()
		fullName, _ := vo2.NewFullName("Alexey Rachkov")
		email, _ := vo2.NewEmail("alexey@gmail.com")
		username, _ := vo2.NewUsername("alexey")

		e := NewAccountCreatedEvent(id, *fullName, *email, *username)

		So(e.AccountID(), ShouldEqual, id)
		So(e.FullName(), ShouldEqual, *fullName)
		So(e.Email(), ShouldEqual, *email)
		So(e.Username(), ShouldEqual, *username)
		So(e.GetEventName(), ShouldEqual, "EVENT.ACCOUNT_CREATED")
	})
}

package domainevent

import (
	vo "account-management-service/internal/valueobject"
	"github.com/google/uuid"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAccountCreatedEvent(t *testing.T) {
	Convey("Test account created event", t, func() {
		id, _ := uuid.NewUUID()
		fullName, _ := vo.NewFullName("Alexey Rachkov")
		email, _ := vo.NewEmail("alexey@gmail.com")
		username, _ := vo.NewUsername("alexey")

		e := NewAccountCreatedEvent(id, *fullName, *email, *username)

		So(e.AccountID(), ShouldEqual, id)
		So(e.FullName(), ShouldEqual, *fullName)
		So(e.Email(), ShouldEqual, *email)
		So(e.Username(), ShouldEqual, *username)
		So(e.GetEventName(), ShouldEqual, "EVENT.ACCOUNT_CREATED")
	})
}

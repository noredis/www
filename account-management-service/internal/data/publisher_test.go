package data

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFakePublisher(t *testing.T) {
	Convey("Test fake publisher", t, func() {
		Convey("Write and read", func() {
			p := NewFakePublisher()
			err := p.Publish("asd")

			So(err, ShouldBeNil)

			var messages []any
			messages = append(messages, "asd")

			So(p.Messages(), ShouldResemble, messages)
		})
	})
}

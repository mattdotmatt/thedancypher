package engine

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCypher(t *testing.T) {

	Convey("Given I am encrypting a message", t, func() {

		Convey("When the message is hello", func() {

			result := Encrypt("hello")

			Convey("Then the result should be xcxcxc", func() {

				So(result, ShouldEqual, "")

			})
		})
	})
}

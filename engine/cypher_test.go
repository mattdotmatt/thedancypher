package engine

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCypher(t *testing.T) {

	Convey("Given I am encrypting a message", t, func() {

		Convey("When the message is xyz", func() {

			result := Encrypt("xyz")

			Convey("Then the result should be xyz", func() {

				So(string(result), ShouldEqual, "ҖҖΞ-")

			})
		})
	})
}

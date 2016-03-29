package engine

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStep1(t *testing.T) {

	Convey("Given I am encrypting a message using step 1", t, func() {

		Convey("When the message is hEllO my Name iS dan", func() {

			result := Step1("hEllO my Name iS dan")

			Convey("Then the result should be EllOha ym meNa iS anda", func() {

				So(result, ShouldEqual, "EllOha ym meNa iS anda")

			})
		})
	})
}

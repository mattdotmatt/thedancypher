package engine

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStep1(t *testing.T) {

	Convey("Given I am encrypting a message using step 1", t, func() {

		Convey("When the message is hello my name is dan", func() {

			result := Step1("hello my name is dan")

			Convey("Then the result should be elloha ym mena is anda", func() {

				So(result, ShouldEqual, "elloha ym mena is anda")

			})
		})
	})
}

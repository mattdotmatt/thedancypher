package engine

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStep3(t *testing.T) {

	Convey("Given I am encrypting a message using step 3", t, func() {

		Convey("When the message is abCdE bAd", func() {

			result := Step3("abCdE bAd")

			Convey("Then the result should be -|\\/Д,|-/", func() {

				So(result, ShouldEqual, "-|\\/Д,|-/")

			})
		})
	})
}

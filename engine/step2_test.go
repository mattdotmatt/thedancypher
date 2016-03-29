package engine

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStep2(t *testing.T) {

	Convey("Given I am encrypting a message using step 2", t, func() {

		Convey("When the message is Elloha ym mena is anda", func() {

			result := Step2("Elloha ym mena is anda")

			Convey("Then the result should be RYYBUN LZ ZRAN VF NAQN", func() {

				So(result, ShouldEqual, "RYYBUN LZ ZRAN VF NAQN")

			})
		})
	})
}

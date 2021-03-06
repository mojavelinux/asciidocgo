package asciidocgo

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAsciidocgo(t *testing.T) {
	Load(nil)
	Convey("Asciidocgo load() takes a string and return a Document", t, func() {
		Convey("A empty string must returns a nil Document", func() {
			So(LoadString(""), ShouldBeNil)
		})
	})
	Convey("Asciidocgo load() takes a array and return a Document", t, func() {
		Convey("A empty array of strings must returns a nil Document", func() {
			So(LoadStrings(), ShouldBeNil)
		})
	})
	Convey("Asciidocgo load() takes a Reader and return a Document", t, func() {
		Convey("A nil Reader must returns a nil Document", func() {
			So(Load(nil), ShouldBeNil)
		})
	})
}

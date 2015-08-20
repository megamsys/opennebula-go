package flavor

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetTemplate(t *testing.T) {
	Convey("GetTemplate", t, func() {
		endpoint := "http://localhost:2633/RPC2"
		key := "oneadmin:RaifZuewjoc4"
		flav := FlavorOpts{TemplateId: 3}
		_, error := flav.GetTemplate(endpoint, key)
		So(error, ShouldBeNil)
	})
}

/*
func TestGetAllTemplates(t *testing.T) {
	Convey("GetAllTemplates", t, func() {

		endpoint := "http://localhost:2633/RPC2"
		key := "oneadmin:RaifZuewjoc4"
		flav := FlavorOpts{}
		_, error := flav.GetAllTemplates(endpoint, key)
		So(error, ShouldBeNil)

	})
}
*/
